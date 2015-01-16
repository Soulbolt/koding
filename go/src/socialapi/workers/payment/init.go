package payment

import (
	"fmt"
	"socialapi/config"
	"socialapi/workers/payment/paymentmodels"
	"socialapi/workers/payment/paypal"
	"socialapi/workers/payment/stripe"
	"time"

	kiteConfig "github.com/koding/kite/config"

	"github.com/koding/bongo"
	"github.com/koding/kite"
)

func Initialize(conf *config.Config) {
	stripe.InitializeClientKey(conf.Stripe.SecretToken)
	paypal.InitializeClientKey(conf.Paypal)

	KiteClient = initializeKiteClient(conf.Kloud.SecretKey, conf.Kloud.Address)

	go func() {
		err := stripe.CreateDefaultPlans()
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
	}()

	go InitCheckers()
}

func InitCheckers() error {
	err := CheckForLeakedSubscriptions()
	if err != nil {
		Log.Error("Error checking for leaked subscriptions %v", err)
	}

	return err
}

func CheckForLeakedSubscriptions() error {
	thirtyDaysAgo := time.Now().Add(-30 * 24 * time.Hour)

	subscription := paymentmodels.NewSubscription()
	subscriptions, err := subscription.ByCanceledAtGte(thirtyDaysAgo)
	if err != nil {
		if err == bongo.RecordNotFound {
			return nil
		}

		return err
	}

	subscriptionIds := []int64{}

	for _, subscription := range subscriptions {
		subscriptionIds = append(subscriptionIds, subscription.Id)
	}

	if len(subscriptions) > 0 {
		Log.Error(
			"%v subscriptions have been expired for more than 30 days. %v",
			len(subscriptionIds), subscriptionIds,
		)
	}

	return nil
}

func initializeKiteClient(kloudSecretKey, kloudAddr string) *kite.Client {
	var err error

	// create new kite
	k := kite.New(WorkerName, WorkerVersion)
	config, err := kiteConfig.Get()
	if err != nil {
		Log.Error("%s. initializing kite failed", err.Error())
		return nil
	}

	// set skeleton config
	k.Config = config

	// create a new connection to the cloud
	kiteClient := k.NewClient(kloudAddr)
	kiteClient.Auth = &kite.Auth{
		Type: "kloudctl",
		Key:  kloudSecretKey,
	}

	// dial the kloud address
	if err := kiteClient.DialTimeout(time.Second * 10); err != nil {
		Log.Error("%s. Is kloud/kontrol running?", err.Error())
		return nil
	}

	Log.Info("Connected to klient: %s", kloudAddr)

	return kiteClient
}
