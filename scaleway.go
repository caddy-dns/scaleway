package scaleway

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/libdns/scaleway"
)

// Provider lets Caddy read and manipulate DNS records hosted by this DNS provider.
type Provider struct{ *scaleway.Provider }

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.scaleway",
		New: func() caddy.Module { return &Provider{new(scaleway.Provider)} },
	}
}

// Provision sets up the module. Implements caddy.Provisioner.
func (p *Provider) Provision(ctx caddy.Context) error {
	p.Provider.SecretKey = caddy.NewReplacer().ReplaceAll(p.Provider.SecretKey, "")
	p.Provider.OrganizationID = caddy.NewReplacer().ReplaceAll(p.Provider.OrganizationID, "")
	return nil
}

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
//	scaleway {
//	    secret_key string
//	    organization_id string
//	}
//
// **THIS IS JUST AN EXAMPLE AND NEEDS TO BE CUSTOMIZED.**
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "secret_key":
				if p.Provider.SecretKey != "" {
					return d.Err("Secret key already set")
				}
				if d.NextArg() {
					p.Provider.SecretKey = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "organization_id":
				if p.Provider.OrganizationID != "" {
					return d.Err("Organization ID already set")
				}
				if d.NextArg() {
					p.Provider.OrganizationID = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}
	if p.Provider.SecretKey == "" {
		return d.Err("missing Secret key")
	}
	if p.Provider.OrganizationID == "" {
		return d.Err("missing Organization ID")
	}
	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
