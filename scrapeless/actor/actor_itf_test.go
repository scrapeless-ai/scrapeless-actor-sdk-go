package actor

import (
	"context"
	"fmt"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/browser"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/captcha"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/proxies"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/storage/kv"
	"testing"
)

func TestActor1(t *testing.T) {
	ac := New(WithBrowser("http"))
	create, err := ac.Browser.Create(context.Background(), browser.Actor{
		Input:        browser.Input{SessionTtl: "180"},
		ProxyCountry: "US",
	})
	defer ac.Close()
	t.Log(create)
	t.Log(err)
}

func TestActorProxy(t *testing.T) {
	//httpProxy := proxies.NewPHttp()
	//create, err := httpProxy.Proxy(context.Background(), proxies.ProxyActor{
	//	Country:         "US",
	//	SessionDuration: 121,
	//	SessionId:       "sessionId",
	//})
	//t.Log(create)
	//t.Error(err)
	ac := New(WithProxy())
	create, err := ac.Proxy.Proxy(context.Background(), proxies.ProxyActor{
		Country:         "US",
		SessionDuration: 121,
		SessionId:       "sessionId",
	})
	t.Log(create)
	t.Error(err)
}

func TestActor_captcha(t *testing.T) {
	ac := New(WithCaptcha())
	create, err := ac.Captcha.Create(context.Background(), &captcha.CaptchaSolverReq{
		Actor: "captcha.recaptcha",
		Input: captcha.Input{
			Version: captcha.RecaptchaVersionV2,
			PageURL: "https://venue.cityline.com",
			SiteKey: "6Le_J04UAAAAAIAfpxnuKMbLjH7ISXlMUzlIYwVw",
		},
		Proxy: captcha.ProxyInfo{
			Country: "US",
		},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(create)
}

func TestActor_captchaResult(t *testing.T) {
	ac := New(WithCaptcha())
	create, err := ac.Captcha.ResultGet(context.Background(), &captcha.CaptchaSolverReq{
		TaskId: "task_id",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(create)
}

func TestActor_captchaSolve(t *testing.T) {
	ac := New(WithCaptcha())
	create, err := ac.Captcha.Solver(context.Background(), &captcha.CaptchaSolverReq{
		Actor: "captcha.recaptcha",
		Input: captcha.Input{
			Version: captcha.RecaptchaVersionV2,
			PageURL: "https://venue.cityline.com",
			SiteKey: "6Le_J04UAAAAAIAfpxnuKMbLjH7ISXlMUzlIYwVw",
		},
		Proxy: captcha.ProxyInfo{
			Country: "US",
		},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(create)
}

func TestActor_kvListNamespace(t *testing.T) {
	ac := New(WithStorage())
	namespaces, err := ac.Storage.GetKv().ListNamespaces(context.TODO(), 1, 10, false)
	fmt.Println(err)
	fmt.Printf("%+v", namespaces)
}

func TestActor_kvCreatNamespace(t *testing.T) {
	ac := New(WithStorage())
	namespaces, _, err := ac.Storage.GetKv().CreateNamespace(context.TODO(), "test")
	fmt.Println(err)
	fmt.Printf("%+v", namespaces)
}

func TestActor_kvGetNamespace(t *testing.T) {
	ac := New(WithStorage())
	namespace, err := ac.Storage.GetKv().GetNamespace(context.Background(), "")
	fmt.Println(err)
	fmt.Printf("%+v", namespace)
}

func TestActor_kvDelNamespace(t *testing.T) {
	ac := New(WithStorage())
	namespace, err := ac.Storage.GetKv().DelNamespace(context.Background())
	fmt.Println(err)
	fmt.Printf("%+v", namespace)
}

func TestActor_kvRenameNamespace(t *testing.T) {
	ac := New(WithStorage())
	namespace, _, err := ac.Storage.GetKv().RenameNamespace(context.Background(), "test-cy-2")
	fmt.Println(err)
	fmt.Printf("%+v", namespace)
}

func TestActor_ListKeys(t *testing.T) {
	ac := New(WithStorage())
	keys, err := ac.Storage.GetKv().ListKeys(context.Background(), 0, 0)
	fmt.Println(err)
	fmt.Printf("%+v", keys)
}

func TestKvActor_DelValue(t *testing.T) {
	ac := New(WithStorage())
	fmt.Println(ac.Storage.GetKv().DelValue(context.Background(), "C"))
}

func TestKvActor_BulkSetValue(t *testing.T) {
	ac := New(WithStorage())
	count, err := ac.Storage.GetKv().BulkSetValue(context.Background(), []kv.BulkItem{
		{
			Key:        "testKey111",
			Value:      "dmFsdWUy",
			Expiration: 3600,
		},
	})
	fmt.Println(err)
	fmt.Println(count)
}

func TestKvActor_BulkDelValue(t *testing.T) {
	ac := New(WithStorage())
	value, err := ac.Storage.GetKv().BulkDelValue(context.Background(), []string{"testKey2", "testKey3"})
	fmt.Println(err)
	fmt.Printf("%+v", value)
}
