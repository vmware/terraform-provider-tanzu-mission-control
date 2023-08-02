# Create explicit cluster proxy credential
resource "tanzu-mission-control_credential" "explicit_proxy_cred" {
  name = "explicit_proxy_cred"

  meta {
    description = "explicit proxy credential"
    labels = {
      "key1" : "value1",
    }
    annotations = {
      "proxyType" : "explicit",
      "httpProxy" : "http://sfsdf.com:123",
      "httpsProxy" : "http://sfsdf.com:123",
      "noProxyList" : "http://noproxy.com,http://something.com"
    }
  }

  spec {
    capability = "PROXY_CONFIG"
    provider   = "GENERIC_KEY_VALUE"
    data {
      key_value {
        data = {
          "httpUserName"  = "username"
          "httpPassword"  = "password"
          "httpsUserName" = "username"
          "httpsPassword" = "password"
          "proxyCABundle" = "-----BEGIN CERTIFICATE-----\n Encoded string for encryption of data\n ----END CERTIFICATE----" # chain of certificate is supported in CRT format
        }
      }
    }
  }
}

# Create transparent cluster proxy credential
resource "tanzu-mission-control_credential" "transparent_proxy_cred" {
  name = "transparent_proxy_cred"

  meta {
    description = "transparent proxy credential"
    labels = {
      "key1" : "value1",
    }
    annotations = {
      "proxyType" : "transparent",
      "noProxyList" : "http://noproxy.com,http://something.com"
    }
  }

  spec {
    capability = "PROXY_CONFIG"
    provider   = "GENERIC_KEY_VALUE"
    data {
      key_value {
        data = {
          "proxyCABundle" = "-----BEGIN CERTIFICATE-----\n Encoded string for encryption of data\n ----END CERTIFICATE----" # chain of certificate is supported in CRT format
        }
      }
    }
  }
}
