# Cash Drawer Banner

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## Update cash drawer banner

```
PATCH /cash_drawer_banner
```

Update a cash drawer banner by ID.

**Response 204** — no content.

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/cash_drawer_banner.json
```

---

### UpdateCashDrawerBannerRequest schema {#updatecashdrawerbannerrequest-schema}

| Field | Type | Required |
|---|---|---|
| `dismissed` | boolean | Yes |

