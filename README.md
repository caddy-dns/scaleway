Scaleway module for Caddy
===========================

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records with Scaleway.

## Caddy module name

```
dns.providers.scaleway
```

## Config examples

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/) like so:

```json
{
	"module": "acme",
	"challenges": {
		"dns": {
			"provider": {
				"name": "scaleway",
				"secret_key": "{env.SCW_ACCESS_KEY}",
				"organization_id": "{env.SCW_ORGANIZATION_ID}"
			}
		}
	}
}
```

or with the Caddyfile:

```
tls {
	dns scaleway {
		secret_key {env.SCW_ACCESS_KEY}
		organization_id {env.SCW_ORGANIZATION_ID}
	}
}
```


## Authenticating

See [the associated README in the libdns package](https://github.com/libdns/scaleway) for important information about credentials.
