# Create cluster proxy credential
resource "tanzu-mission-control_credential" "proxy_cred" {
  name = "proxy_cred"

  meta {
    description = "proxy credential"
    labels = {
      "key1" : "value1",
    }
    annotations = {
      "httpProxy"  :"http://sfsdf.com:123",
      "httpsProxy" :"http://sfsdf.com:123",
      "noProxyList":"http://noproxy.com,http://something.com"
    }
  }

  spec {
    capability = "PROXY_CONFIG"
    provider = "GENERIC_KEY_VALUE"
    data {
      key_value{
        data = {
          "httpUserName"  = "username"
          "httpPassword"  = "password"
          "httpsUserName" = "username"
          "httpsPassword" = "password"
          "proxyCABundle" = "cabundle"
        }
      }
    }
  }
}
