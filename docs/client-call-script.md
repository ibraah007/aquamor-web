# Client call script: what to ask for, in plain language

Use this when the client calls instead of replying in writing. It explains every
item in [`client-information-needed.md`](client-information-needed.md) in the
client's own terms, with the questions he is likely to ask and short answers.
The technical version is in [`go-live-runbook.md`](go-live-runbook.md); this file
is the talk track.

**Aim of the call:** get four things — approval, domain ownership, the business
details, and a decision on hosting. Nothing else is blocking launch.

**If he wants one sentence:** "Register the domain in the business name, buy a
simple monthly hosting account, send me the logins, and approve the site — I do
the rest."

---

## Before the call

Have these six facts ready:

1. Domain: about **Ksh 1,000–2,000 per year** for `.co.ke`, registered in the
   business name, auto-renew on.
2. Hosting: about **Ksh 1,500–3,000 per month**, a simple plan — not an
   e-commerce or website-builder package.
3. Documents: the client sends them to the registrar, not to us. We only need the
   exact business name, email, phone and address.
4. DNS: two records, or a delegated login. No password sharing.
5. Preview vs live: the preview saves nothing; the live site saves orders.
6. Launch: **10 December 2026**, about 1–2 hours of work once the domain and
   hosting exist.

Preview link: <https://ibraah007.github.io/aquamor-web/>

---

## Part 1 — Approval

**What we are asking for:** one written line saying the site is approved.

**Why:** it freezes the content. Once we deploy, we do not want "actually, change
the delivery areas" arriving halfway through launch. It is not a contract, just a
checkpoint.

**If he asks:**

- *"Can I change things later?"* — Yes. Text, prices and photos are easy, usually
  the same day. Structural redesigns take longer.
- *"Can I just say it on the phone?"* — He can, but ask him to send it as a
  message too, so there is a record of what was approved.

---

## Part 2 — Domain

**In his language:** the domain is the shop's name on the street. It is rented
yearly, not bought once.

**Why `aquamorwater.co.ke`:** it matches the business name, `.co.ke` builds local
trust, and the site already uses it.

**Why it must be in the business name:** if it is registered in a personal name,
that person legally owns Aquamor's address. If the working relationship ever ends,
that is a hostage situation. In the business name, Aquamor owns it, pays the
renewal and keeps it whoever built the site.

**Why we need DNS access:** DNS is the forwarding address — it tells the internet
that `aquamorwater.co.ke` lives on our server. That is two records. We do not need
his registrar password: the registrar can give us a limited "delegated user"
login, or he types two lines we send him.

**Why the screenshot:** proof the domain is active, proof of the expiry date, and
proof the name is spelled correctly. "I registered it" with a typo costs a week to
fix.

**Why auto-renew:** so the site does not disappear in 12 months because a renewal
email went to an old inbox.

**What to ask him to send:**

- A screenshot of the registrar dashboard showing the domain, **Active** status,
  expiry date and registrant name.
- Which registrar he used, and who holds the login.
- The registrant email and phone — one he actually reads.
- Confirmation that auto-renew is on.
- DNS access: either add us as a delegated user with DNS rights, or he pastes two
  records we send.

**If he asks:**

- *"Why can't you just register it?"* — It must be in the business name, and the
  registrar needs his certificate or KRA PIN and his payment. Registering it
  ourselves would put it in our name, which is what we are avoiding.
- *"How much?"* — Roughly Ksh 1,000–2,000 per year. It varies by registrar.
- *"Which registrar?"* — Any KeNIC-accredited one, for example Truehost or
  HostPinnacle. Whoever he is comfortable paying.
- *"Can we use `.com` instead?"* — Possible, but `.co.ke` is better for a Kenyan
  water business. Decide now, because it changes links across the site.
- *"Will I lose my email?"* — No. The domain and email are separate; we only point
  the website.

**Watch out:** if he buys hosting in the same registrar checkout, he pays for
something we do not need. Say it plainly: **domain only**.

---

## Part 3 — Documents

**What is happening:** the registrar — not us — needs proof the business exists.
That is a KeNIC rule for a business `.co.ke` domain.

**What he sends the registrar:** the business registration certificate or KRA PIN.

**What he sends us:** only the exact business name spelling, the email, the phone
and the address used on the domain. Nothing sensitive.

**Why the spelling matters:** it goes on the domain record and the site footer.
"AQUAMOR PURIFIED DRINKING WATER" must match everywhere.

**If he asks:**

- *"Do you need my KRA PIN?"* — No. Only the registrar does. Do not send it to us
  and do not send it in a group chat.
- *"Can I register it in my own name with my ID?"* — He can, but then the domain is
  personal property, not the business's. Recommend against it; if he insists, note
  it and move on.
- *"The registrar rejected the business name."* — It happens. Options: register
  personally and transfer later, or ask the registrar which document they will
  accept.

---

## Part 4 — Hosting

**The analogy:** the preview is a **photo of the shop**. People can look, but
nobody can buy anything. Hosting is the **actual shop, open, with a till that
records sales**.

**Why the preview cannot be the live site:** it does not save orders, and there is
no staff page to see incoming orders. It is a demo, not a business.

**Why a "small computer that stays on 24/7":** that is all a server is. A site
that takes orders has to be running every minute, and something has to store the
orders. That is what the monthly fee pays for.

**The three options:**

| Option | What happens | Trade-off |
|---|---|---|
| **A — He buys** (recommended) | He opens an account in Aquamor's name, pays, sends us the login | He owns everything; slightly more admin for him |
| **B — We buy** | We set it up and hand over the login at launch | Fastest for him; ownership transfers at the end, so agree it in writing |
| **C — He already has hosting** | He sends the login, we check it | It may be too cheap or too shared to run the order system; we will tell him |

**Cost:** about **Ksh 1,500–3,000 per month**. Monthly or yearly; yearly is
usually cheaper.

**What he must not buy:** an "e-commerce plan", a "website builder", "unlimited
business hosting" or a "free domain included" bundle. Those are upsells. He needs
the cheapest plan that gives a login and can run a small database. If the hosting
company asks what he needs, he should say: **"the cheapest plan that comes with a
login and can run a small database — not a website builder or e-commerce plan."**

**What "we look after it monthly" means:** security updates, backups, watching
that it is online, and fixing it if it breaks. If he says no, we hand over the
keys and he is responsible — including when it goes down.

**If he asks:**

- *"Can't we just use Facebook and WhatsApp?"* — The site does not replace them.
  It is the shopfront that makes Aquamor look established and gets found on
  Google. Orders still come through WhatsApp.
- *"Why monthly? I thought a website is one-time."* — The building is one-time.
  Rent, electricity and security are monthly. Hosting is the rent.
- *"Can I pay yearly?"* — Yes, and it is usually cheaper. Just keep auto-renew on.
- *"What if it goes down?"* — That is what the monthly management covers. Without
  it, someone has to notice and fix it.

**Security, say it once:** send the hosting password through a password manager or
read it out on the call. Never in a WhatsApp group or a screenshot.

---

## What happens after each item arrives

| He sends | We do | Time |
|---|---|---|
| Approval | Freeze content, final checks | Same day |
| Domain screenshot and access | Set the DNS records, confirm they resolve | About an hour, plus propagation |
| Business name, email, phone | Update the site footer if needed | Minutes |
| Hosting login | Install, deploy, HTTPS, backups, test | 1–2 hours |

Then he opens the site on his phone, we run the smoke test, hand over the staff
order page and the admin key, and monitor for the first week.

---

## Forwardable message

Send this after the call as a written summary. It is deliberately free of
technical words.

```text
Hi [name], here is what I need to take the site live, as we discussed.

1. APPROVAL
   Reply: "The site is approved as it is" — or send your changes now.
   Preview: https://ibraah007.github.io/aquamor-web/

2. DOMAIN (aquamorwater.co.ke)
   Register it in the business name AQUAMOR PURIFIED DRINKING WATER, then send me:
   - a screenshot of the dashboard showing the domain, Active status,
     expiry date and registrant name
   - which registrar you used, and who holds the login
   - the email and phone on the domain (one you actually read)
   - confirmation that auto-renew is ON
   - DNS access: either add me as a delegated user with DNS rights,
     or I will send you two records to paste in yourself
   Do not buy hosting from the registrar, and keep the domain in Aquamor's
   name — I only need technical access.

3. DOCUMENTS
   The registrar will ask you for the business certificate or KRA PIN. That is
   between you and them; I do not need to see them. Just tell me the exact
   business name as written on the certificate, and the email, phone and address
   used on the domain.

4. HOSTING (the place the live site runs)
   The preview is only a demo — it cannot save orders. For the real site we need
   a small hosting account that stays online all the time. Cost is roughly
   Ksh 1,500–3,000 per month. You do NOT need the expensive "business" or
   "e-commerce" package; the basic one is enough.

   Pick one:
   a) You open the hosting account in Aquamor's name and send me the login.
      I set everything up. (Recommended — you own it.)
   b) I buy it and set it up, and give you the login at launch.
   c) You already have hosting — send me the login and I will check it.

   Also tell me:
   - should I look after it every month (updates, backups, keeping it online),
     or hand everything to you after launch?
   - how would you like to pay: your card, or reimburse me?

   Send any password through a password manager or read it to me on a call,
   not in a WhatsApp group.

Once the domain and hosting are ready I deploy, test it on your phone, and hand
over the staff order page. Target: 10 December 2026.
```

---

## After the call

1. Send the written summary above, so both sides have the same list.
2. Log what was agreed in [`maintenance-log.md`](maintenance-log.md).
3. Tick the matching items in [`launch-checklist.md`](launch-checklist.md).
4. When the domain and hosting exist, follow [`go-live-runbook.md`](go-live-runbook.md).

Do not start the domain or DNS work until the approval is in writing.
