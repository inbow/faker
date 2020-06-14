package server

import (
	"net/http"
	"time"

	"github.com/savsgio/atreugo/v11"
)

func (s *Server) OpenRTBNativeMultiBid(ctx *atreugo.RequestCtx) error {
	response := s.NewResponse()
	_, delay, skip := s.RequestValues(ctx.QueryArgs())

	defer func() {
		time.Sleep(time.Duration(delay) * time.Millisecond)

		ctx.Response.Header.Set("Content-Type", "application/json")
		ctx.SetStatusCode(response.StatusCode)

		if response.StatusCode != http.StatusNoContent && len(response.Body) > 0 {
			ctx.SetBody(response.Body)
		}
	}()

	if skip {
		response.StatusCode = http.StatusNoContent
		return nil
	}

	response.StatusCode = http.StatusOK
	response.Body = openRTBNativeMultiBid

	return nil
}

var (
	// nolint:lll
	openRTBNativeMultiBid = []byte(`
{
   "id":"e1416851-e170-435b-b857-eb0c7609af81",
   "seatbid":[
      {
         "bid":[
            {
               "id":"1native:28270c2d-6e6c-4c44-a820-041e2a49031c",
               "impid":"28270c2d-6e6c-4c44-a820-041e2a49031c",
               "price":0.05,
               "adm":"{\"native\":{\"link\":{\"url\":\"https:\\\/\\\/eu-adsrv.rtbsuperhub.com\\\/click\\\/?subPublisher=native:xhamster.com&zone=native:xhamster.com&adformat=native&auctionid=e1416851-e170-435b-b857-eb0c7609af81&uniqueid=b1977da102a3f9471c0cffc0726bd068&name=trafficstars-native_native_usa_tablet_xHamster&campaign={campaign}&width=300&height=250&newservice=true&cmsid=landing--mlp6005--landing--cm8001&tpcampid=3c84f5e5-06b7-4ec4-a58c-da46f112c2b6&imp_tagid=2811&ba=2158b725-92cd-4c22-a6dd-56f9110ead3f&uid=b8095d41-9d78-4f6d-ad11-1d211904846a&campaign_lp=1:landing--mlp6005--landing--cm8001&product=dateyouweb\"},\"assets\":[{\"id\":3,\"required\":1,\"img\":{\"url\":\"https:\\\/\\\/bmedia.justservingfiles.net\\\/a9b2bc09-19eb-4473-86db-a6e69e8914bd.jpg\",\"w\":300,\"h\":300}},{\"id\":1,\"title\":{\"text\":\"This free sex site for adults has taken the net by storm!\"}},{\"id\":2,\"data\":{\"value\":\"myDates\"}}],\"imptrackers\":[],\"eventtrackers\":[{\"event\":1,\"method\":1,\"url\":\"https:\\\/\\\/eu-adsrv.rtbsuperhub.com\\\/callback?method=1&landingpage=22158b725-92cd-4c22-a6dd-56f9110ead3f&b8095d41-9d78-4f6d-ad11-1d211904846a&0\"},{\"event\":1,\"method\":2,\"url\":\"https:\\\/\\\/eu-adsrv.rtbsuperhub.com\\\/callback?method=2&landingpage=22158b725-92cd-4c22-a6dd-56f9110ead3f&b8095d41-9d78-4f6d-ad11-1d211904846a&0\"}]}}",
               "nurl":"http:\/\/tp-rtb-us.us-east-2.elasticbeanstalk.com\/callback\/?network=trafficstars-native&adformat=native&uniqueid=1&type=0&zone=xhamster.com&account=trafficpartner_rtb&name=trafficstars-native_native_usa_tablet_xHamster&device_geo_country=USA&auction_id=e1416851-e170-435b-b857-eb0c7609af81&auction_bid_id=${AUCTION_BID_ID}&auction_imp_id=${AUCTION_IMP_ID}&auction_seat_id=${AUCTION_SEAT_ID}&auction_ad_id=${AUCTION_AD_ID}&auction_price=${AUCTION_PRICE}&auction_currency=${AUCTION_CURRENCY}&campaign=trafficstars-native-native-ios&br=2&bw=23&bannerid=2158b725-92cd-4c22-a6dd-56f9110ead3f&campaign_lp=1:landing--mlp6005--landing--cm8001&imp_tagid=2811&udid=b8095d41-9d78-4f6d-ad11-1d211904846a&exploration_banner=bestCPA&exploration_landingpage=lpranking&exploration_cpa=exoloration&exploration_ctr=normal&abtests=0&product=dateyouweb&cookie_fallback=9af9d8327be8edc01e49ba99b023e00a&banner_ecpa=18.672799142857&landingpage_ecpa=1&srtbid=0&device_os=ios"
            }
         ]
      },
      {
         "bid":[
            {
               "id":"2native:28270c2d-6e6c-4c44-a820-041e2a49031c",
               "impid":"28270c2d-6e6c-4c44-a820-041e2a49031c",
               "price":0.05,
               "adm":"{\"native\":{\"link\":{\"url\":\"https:\\\/\\\/eu-adsrv.rtbsuperhub.com\\\/click\\\/?subPublisher=native:xhamster.com&zone=native:xhamster.com&adformat=native&auctionid=e1416851-e170-435b-b857-eb0c7609af81-native2&uniqueid=b1977da102a3f9471c0cffc0726bd068&name=trafficstars-native_native_usa_tablet_xHamster&campaign={campaign}&width=300&height=250&newservice=true&cmsid=landing--mlp6005--landing--cm8001&tpcampid=3c84f5e5-06b7-4ec4-a58c-da46f112c2b6&imp_tagid=2811&ba=cb22ee1e-b12d-4394-89a4-5150d5dcd976&uid=b8095d41-9d78-4f6d-ad11-1d211904846a&campaign_lp=1:landing--mlp6005--landing--cm8001&product=dateyouweb\"},\"assets\":[{\"id\":3,\"required\":1,\"img\":{\"url\":\"https:\\\/\\\/bmedia.justservingfiles.net\\\/6e8a1868-c4c8-4215-84e9-b7b251525b3f.jpg\",\"w\":300,\"h\":300}},{\"id\":1,\"title\":{\"text\":\"Meet girls near you for sex now!\"}},{\"id\":2,\"data\":{\"value\":\"myDates\"}}],\"imptrackers\":[],\"eventtrackers\":[{\"event\":1,\"method\":1,\"url\":\"https:\\\/\\\/eu-adsrv.rtbsuperhub.com\\\/callback?method=1&landingpage=2cb22ee1e-b12d-4394-89a4-5150d5dcd976&b8095d41-9d78-4f6d-ad11-1d211904846a&0\"},{\"event\":1,\"method\":2,\"url\":\"https:\\\/\\\/eu-adsrv.rtbsuperhub.com\\\/callback?method=2&landingpage=2cb22ee1e-b12d-4394-89a4-5150d5dcd976&b8095d41-9d78-4f6d-ad11-1d211904846a&0\"}]}}",
               "nurl":"http:\/\/tp-rtb-us.us-east-2.elasticbeanstalk.com\/callback\/?network=trafficstars-native&adformat=native&uniqueid=2&type=0&zone=xhamster.com&account=trafficpartner_rtb&name=trafficstars-native_native_usa_tablet_xHamster&device_geo_country=USA&auction_id=e1416851-e170-435b-b857-eb0c7609af81-native2&auction_bid_id=${AUCTION_BID_ID}&auction_imp_id=${AUCTION_IMP_ID}&auction_seat_id=${AUCTION_SEAT_ID}&auction_ad_id=${AUCTION_AD_ID}&auction_price=${AUCTION_PRICE}&auction_currency=${AUCTION_CURRENCY}&campaign=trafficstars-native-native-ios&br=2&bw=23&bannerid=cb22ee1e-b12d-4394-89a4-5150d5dcd976&campaign_lp=1:landing--mlp6005--landing--cm8001&imp_tagid=2811&udid=b8095d41-9d78-4f6d-ad11-1d211904846a&exploration_banner=bestCPA&exploration_landingpage=lpranking&exploration_cpa=exoloration&exploration_ctr=normal&abtests=0&product=dateyouweb&cookie_fallback=9af9d8327be8edc01e49ba99b023e00a&banner_ecpa=18.672799142857&landingpage_ecpa=1&srtbid=0&device_os=ios"
            }
         ]
      },
      {
         "bid":[
            {
               "id":"3native:28270c2d-6e6c-4c44-a820-041e2a49031c",
               "impid":"28270c2d-6e6c-4c44-a820-041e2a49031c",
               "price":0.05,
               "adm":"{\"native\":{\"link\":{\"url\":\"https:\\\/\\\/eu-adsrv.rtbsuperhub.com\\\/click\\\/?subPublisher=native:xhamster.com&zone=native:xhamster.com&adformat=native&auctionid=e1416851-e170-435b-b857-eb0c7609af81-native3&uniqueid=b1977da102a3f9471c0cffc0726bd068&name=trafficstars-native_native_usa_tablet_xHamster&campaign={campaign}&width=300&height=250&newservice=true&cmsid=landing--mlp6005--landing--cm8001&tpcampid=3c84f5e5-06b7-4ec4-a58c-da46f112c2b6&imp_tagid=2811&ba=14cb2478-a2e8-432a-86f4-77f19f47b5f9&uid=b8095d41-9d78-4f6d-ad11-1d211904846a&campaign_lp=1:landing--mlp6005--landing--cm8001&product=dateyouweb\"},\"assets\":[{\"id\":3,\"required\":1,\"img\":{\"url\":\"https:\\\/\\\/bmedia.justservingfiles.net\\\/9ae90c10-87ed-46be-93a7-ba8feb72ab9c.jpg\",\"w\":300,\"h\":300}},{\"id\":1,\"title\":{\"text\":\"See more pictures! Meet girls near you for sex now!\"}},{\"id\":2,\"data\":{\"value\":\"myDates\"}}],\"imptrackers\":[],\"eventtrackers\":[{\"event\":1,\"method\":1,\"url\":\"https:\\\/\\\/eu-adsrv.rtbsuperhub.com\\\/callback?method=1&landingpage=214cb2478-a2e8-432a-86f4-77f19f47b5f9&b8095d41-9d78-4f6d-ad11-1d211904846a&0\"},{\"event\":1,\"method\":2,\"url\":\"https:\\\/\\\/eu-adsrv.rtbsuperhub.com\\\/callback?method=2&landingpage=214cb2478-a2e8-432a-86f4-77f19f47b5f9&b8095d41-9d78-4f6d-ad11-1d211904846a&0\"}]}}",
               "nurl":"http:\/\/tp-rtb-us.us-east-2.elasticbeanstalk.com\/callback\/?network=trafficstars-native&adformat=native&uniqueid=3&type=0&zone=xhamster.com&account=trafficpartner_rtb&name=trafficstars-native_native_usa_tablet_xHamster&device_geo_country=USA&auction_id=e1416851-e170-435b-b857-eb0c7609af81-native3&auction_bid_id=${AUCTION_BID_ID}&auction_imp_id=${AUCTION_IMP_ID}&auction_seat_id=${AUCTION_SEAT_ID}&auction_ad_id=${AUCTION_AD_ID}&auction_price=${AUCTION_PRICE}&auction_currency=${AUCTION_CURRENCY}&campaign=trafficstars-native-native-ios&br=2&bw=23&bannerid=14cb2478-a2e8-432a-86f4-77f19f47b5f9&campaign_lp=1:landing--mlp6005--landing--cm8001&imp_tagid=2811&udid=b8095d41-9d78-4f6d-ad11-1d211904846a&exploration_banner=bestCPA&exploration_landingpage=lpranking&exploration_cpa=exoloration&exploration_ctr=normal&abtests=0&product=dateyouweb&cookie_fallback=9af9d8327be8edc01e49ba99b023e00a&banner_ecpa=18.672799142857&landingpage_ecpa=1&srtbid=0&device_os=ios"
            }
         ]
      },
      {
         "bid":[
            {
               "id":"4native:28270c2d-6e6c-4c44-a820-041e2a49031c",
               "impid":"28270c2d-6e6c-4c44-a820-041e2a49031c",
               "price":0.05,
               "adm":"{\"native\":{\"link\":{\"url\":\"https:\\\/\\\/eu-adsrv.rtbsuperhub.com\\\/click\\\/?subPublisher=native:xhamster.com&zone=native:xhamster.com&adformat=native&auctionid=e1416851-e170-435b-b857-eb0c7609af81-native4&uniqueid=b1977da102a3f9471c0cffc0726bd068&name=trafficstars-native_native_usa_tablet_xHamster&campaign={campaign}&width=300&height=250&newservice=true&cmsid=landing--mlp6005--landing--cm8001&tpcampid=3c84f5e5-06b7-4ec4-a58c-da46f112c2b6&imp_tagid=2811&ba=8495a7a2-8b0d-4fdf-8283-977c3bd7033f&uid=b8095d41-9d78-4f6d-ad11-1d211904846a&campaign_lp=1:landing--mlp6005--landing--cm8001&product=dateyouweb\"},\"assets\":[{\"id\":3,\"required\":1,\"img\":{\"url\":\"https:\\\/\\\/bmedia.justservingfiles.net\\\/876894fb-679a-4196-97a6-92aec41e2209.jpg\",\"w\":300,\"h\":300}},{\"id\":1,\"title\":{\"text\":\"Discrete sex community\"}},{\"id\":2,\"data\":{\"value\":\"myDates\"}}],\"imptrackers\":[],\"eventtrackers\":[{\"event\":1,\"method\":1,\"url\":\"https:\\\/\\\/eu-adsrv.rtbsuperhub.com\\\/callback?method=1&landingpage=28495a7a2-8b0d-4fdf-8283-977c3bd7033f&b8095d41-9d78-4f6d-ad11-1d211904846a&0\"},{\"event\":1,\"method\":2,\"url\":\"https:\\\/\\\/eu-adsrv.rtbsuperhub.com\\\/callback?method=2&landingpage=28495a7a2-8b0d-4fdf-8283-977c3bd7033f&b8095d41-9d78-4f6d-ad11-1d211904846a&0\"}]}}",
               "nurl":"http:\/\/tp-rtb-us.us-east-2.elasticbeanstalk.com\/callback\/?network=trafficstars-native&adformat=native&uniqueid=4&type=0&zone=xhamster.com&account=trafficpartner_rtb&name=trafficstars-native_native_usa_tablet_xHamster&device_geo_country=USA&auction_id=e1416851-e170-435b-b857-eb0c7609af81-native4&auction_bid_id=${AUCTION_BID_ID}&auction_imp_id=${AUCTION_IMP_ID}&auction_seat_id=${AUCTION_SEAT_ID}&auction_ad_id=${AUCTION_AD_ID}&auction_price=${AUCTION_PRICE}&auction_currency=${AUCTION_CURRENCY}&campaign=trafficstars-native-native-ios&br=2&bw=23&bannerid=8495a7a2-8b0d-4fdf-8283-977c3bd7033f&campaign_lp=1:landing--mlp6005--landing--cm8001&imp_tagid=2811&udid=b8095d41-9d78-4f6d-ad11-1d211904846a&exploration_banner=bestCPA&exploration_landingpage=lpranking&exploration_cpa=exoloration&exploration_ctr=normal&abtests=0&product=dateyouweb&cookie_fallback=9af9d8327be8edc01e49ba99b023e00a&banner_ecpa=18.672799142857&landingpage_ecpa=1&srtbid=0&device_os=ios"
            }
         ]
      }
   ]
}
`)
)
