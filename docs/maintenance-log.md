# Maintenance log

A dated record of client decisions, requests and site changes. Add an entry whenever the client confirms a fact, asks for a change, or we make a decision that affects the site.

Keep it short. Link to the issue, PR or file when there is one.

---

## 13 November 2026

### Logo: settled

Decision: use the logo that is already on the site. The header mark (`icon.jpg` / `icon.webp`), `favicon-32.png`, `apple-touch-icon.png` and `icon-512.png` are the approved versions. No original logo file is needed from the client.

Still open from the same item: the official business name and tagline (the client's material has shown more than one version). The site currently uses "Aquamor Purified Drinking Water" and the existing labels.

### Requested from the client (WhatsApp)

Asked for:

- KEBS permit number and a photo of the certificate, if they have one;
- the free-delivery rules (areas and any minimum order).

Message sent:

> Helllo, do you have a KEBS permit?? if yes please send me the permit number and the photo of the certificate and maybe the delivery rules

**Do not add a KEBS number, certification claim, delivery area or minimum order until the client replies and the details are checked.** When the reply arrives:

1. update `docs/status.md` sections 5A and 5D;
2. if a KEBS number is confirmed, add it as a plain fact (no health claims) and keep a copy of the certificate on file;
3. update the delivery copy in `web/index.html` and the Swahili string in `web/js/i18n.js`;
4. rebuild the shop price card only if a price changes;
5. note the change here.

### Client reply (later on 13 November 2026)

Reply received:

> The KEBS is with our supplier we are working on our own next year February we have already gone to the Kebs and we are now in the process,,and also if there is prices for small bottlled water(300ml.500ml,1l,5l)  Branded-bottle pricing and minimum order.
> -  Dispenser-cleaning price.
> [Free-delivery areas and any minimum order
> Oh yes we do them but for now we deal with branded water alone
> Which ranges from 500-800 depending on the type of the bottle

Decisions:

- **KEBS:** the client does not have their own permit or certificate yet; the current certification is with their supplier. Their own application is in process and expected in February next year. Do not add a KEBS number or certification claim. A supplier's certificate is not ours to publish.
- **Prices:** do not publish the Ksh 500–800 range yet. The reply does not say whether it is per bottle or per pack, whether it covers small bottled water or custom-branded bottles, or what the minimum order is.
- **Dispenser cleaning and free delivery:** the client's "yes we do them" is not tied to one question. The cleaning price and the delivery areas/minimum order are still missing.

Follow-up questions sent to the client (13 November 2026); waiting for the reply:

1. Is Ksh 500–800 per bottle or per pack/carton?
2. Which sizes does the Ksh 500–800 range cover?
3. Does that range apply to small bottled water (300ml, 500ml, 1L, 5L) or to custom-branded bottles?
4. What is the minimum order for branded bottles?
5. What is the dispenser-cleaning price?
6. Which areas are free delivery, and is there a minimum order?
7. What is the official business name and tagline?
8. What are the purification steps, if they want them published?

### Earlier decisions this day

- The client has no Instagram account; no Instagram link will be added (issue #5, closed).
- All remaining media was sent and permission to show the people in the photos was confirmed (issue #6, closed).
- There are no real customer testimonials, so the testimonials section stays hidden (issue #8, closed).
- Update model: managed updates. The client sends requests; the maintainer edits, checks and publishes. No self-service CMS in this phase (issue #16).

### Site changes merged

- #11 refill planner, #12 designer presets and round labels, #13 one-tap reorder, #14 how it works and delivery strip, #15 printable price card, #16 handover documents, #17 quality pass, #18 CI checks.
- CI now runs Go build/vet, gofmt and JavaScript syntax on every pull request.

---

## Open items waiting on the client

| Item | Issue | Status |
|---|---|---|
| Small bottled-water prices (300ml, 500ml, 1L, 1.5L) | #4 | Reply ambiguous; confirm unit |
| Branded-bottle pricing and minimum order | #4 | 500–800 quoted; unit/minimum unclear |
| Dispenser-cleaning price | #4 | Waiting |
| Free-delivery areas and minimum order | #4 | Requested 13 Nov; not answered |
| KEBS permit number and a photo of the certificate | #4 | Answered: with supplier; own due Feb |
| Purification steps (only if the client has them) | #4 | Not yet requested |
| Official business name and tagline | #4 | Waiting |
| Facebook page is a business page, not a personal profile | #10 | Waiting |
| Confirm bottled-water sizes (flyer vs client text) | #10 | Waiting |
| Native-speaker Swahili review | #10 | Waiting |
| Regular 20L bottle: is water included? | #10 | Waiting |
| Original cleaning before/after stills | #9 | Waiting |
| More gallery content (shop front, graduation/corporate) | #7 | Waiting |
