package response

import "time"

// PartsMatchResponse represents PartsMatchResponse (converted from a JSON using https://mholt.github.io/json-to-go/)
type PartsMatchResponse struct {
	Class   string `json:"__class__"`
	Msec    int    `json:"msec"`
	Request struct {
		Class     string `json:"__class__"`
		ExactOnly bool   `json:"exact_only"`
		Queries   []struct {
			Class     string      `json:"__class__"`
			Brand     interface{} `json:"brand"`
			Limit     int         `json:"limit"`
			Mpn       string      `json:"mpn"`
			MpnOrSku  interface{} `json:"mpn_or_sku"`
			Q         string      `json:"q"`
			Reference interface{} `json:"reference"`
			Seller    interface{} `json:"seller"`
			Sku       interface{} `json:"sku"`
			Start     int         `json:"start"`
		} `json:"queries"`
	} `json:"request"`
	Results []struct {
		Class string      `json:"__class__"`
		Error interface{} `json:"error"`
		Hits  int         `json:"hits"`
		Items []struct {
			Class string `json:"__class__"`
			Brand struct {
				Class       string `json:"__class__"`
				HomepageURL string `json:"homepage_url"`
				Name        string `json:"name"`
				UID         string `json:"uid"`
			} `json:"brand"`
			Manufacturer struct {
				Class       string `json:"__class__"`
				HomepageURL string `json:"homepage_url"`
				Name        string `json:"name"`
				UID         string `json:"uid"`
			} `json:"manufacturer"`
			Mpn         string `json:"mpn"`
			OctopartURL string `json:"octopart_url"`
			Offers      []struct {
				Class  string `json:"__class__"`
				Custom struct {
				} `json:"_custom"`
				NaiveID              string      `json:"_naive_id"`
				EligibleRegion       string      `json:"eligible_region"`
				FactoryLeadDays      int         `json:"factory_lead_days"`
				FactoryOrderMultiple interface{} `json:"factory_order_multiple"`
				InStockQuantity      int         `json:"in_stock_quantity"`
				IsAuthorized         bool        `json:"is_authorized"`
				IsRealtime           bool        `json:"is_realtime"`
				LastUpdated          time.Time   `json:"last_updated"`
				Moq                  int         `json:"moq"`
				MultipackQuantity    interface{} `json:"multipack_quantity"`
				OctopartRfqURL       interface{} `json:"octopart_rfq_url"`
				OnOrderEta           interface{} `json:"on_order_eta"`
				OnOrderQuantity      interface{} `json:"on_order_quantity"`
				OrderMultiple        interface{} `json:"order_multiple"`
				Packaging            string      `json:"packaging"`
				ProductURL           string      `json:"product_url"`
				Seller               struct {
					Class        string `json:"__class__"`
					DisplayFlag  string `json:"display_flag"`
					HasEcommerce bool   `json:"has_ecommerce"`
					HomepageURL  string `json:"homepage_url"`
					ID           string `json:"id"`
					Name         string `json:"name"`
					UID          string `json:"uid"`
				} `json:"seller"`
				Sku    string `json:"sku"`
				Prices struct {
					EUR [][]interface{} `json:"EUR"`
					SGD [][]interface{} `json:"SGD"`
					USD [][]interface{} `json:"USD"`
					CHF [][]interface{} `json:"CHF"`
				} `json:"prices,omitempty"`
			} `json:"offers"`
			RedirectedUids []string `json:"redirected_uids"`
			UID            string   `json:"uid"`
		} `json:"items"`
		Reference interface{} `json:"reference"`
	} `json:"results"`
}
