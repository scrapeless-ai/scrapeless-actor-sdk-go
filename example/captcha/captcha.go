package main

import (
	"context"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/captcha"
	log "github.com/sirupsen/logrus"
	"time"
)

func main() {
	scrapeless := scrapeless.New(scrapeless.WithCaptcha())
	//Create captcha task
	captchaTaskId, err := scrapeless.Captcha.Create(context.TODO(), &captcha.CaptchaSolverReq{
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
		log.Info(err)
	}
	log.Info(captchaTaskId)
	// Wait for captcha task to be solved
	time.Sleep(time.Second * 20)
	captchaResult, err := scrapeless.Captcha.ResultGet(context.TODO(), &captcha.CaptchaSolverReq{
		TaskId: captchaTaskId,
	})
	if err != nil {
		log.Info(err)
	}
	log.Info(captchaResult)
}
