// Read Tanzu Mission Control cluster proxy credential: fetch details
data "tanzu-mission-control_credential" "test_cred" {
  name = "<credential-name>"
}

// Create/ Delete Tanzu Mission Control cluster proxy credential
resource "tanzu-mission-control_credential" "proxy_cred" {
  name = "<credential-name>"

  meta {
    description = "<description of the credential>"
    labels = {
      "key" : "<value>",
    }
    annotations = {
      "httpProxy" : "<http-proxy-url>",
      "httpsProxy" : "<https-proxy-url>",
      "noProxyList" : "<no-proxy-list>"
    }
  }

  spec {
    capability = "<capability-type>"
    provider   = "<provider>"
    data {
      key_value {
        data = {
          "httpUserName"  = "<username>"
          "httpPassword"  = "<password>"
          "httpsUserName" = "<username>"
          "httpsPassword" = "<password>"
          "proxyCABundle" = "<cabundle>"
        }
      }
    }
  }
}
