# GOJIAS Collections — Data Spec (Plain English)

This document describes the three main "things" the app needs to keep track of, before any code is written. It's meant to be read by a human, not a computer — the Go code will be built from this later.

---

## 1. Customer

A person who wants an outfit made.

- Name
- Contact info (phone and/or email)
- One or more saved measurement profiles (see below) — e.g. "my measurements" and "my daughter's measurements"
- Order history (a list of orders they've placed)

## 2. Tailor

A person who makes the outfits.

- Name
- Contact info
- Location (so customers can find tailors near them)
- Specialties (e.g. "women's dresses," "girls' wear," "traditional styles")
- Portfolio — photos of past finished work
- Rating (average of customer reviews, once they have any)

## 3. Measurement Profile

A saved set of body measurements, so a customer doesn't have to re-enter them every time.

- Belongs to one customer
- A label (e.g. "Me," "Amara — age 8")
- Bust, waist, hips, shoulder width, sleeve length, outfit length (exact fields to be confirmed — this is a starting list)
- Date last updated (measurements change, especially for growing girls)

## 4. Order

The connector between a Customer and a Tailor. This is the record of one specific request.

- Which customer placed it
- Which tailor is making it
- Which measurement profile applies to it (a copy of the numbers at the time of ordering — not a live link, since measurements can change later)
- Fabric choice (print name/color, kitenge or Ankara)
- Style description or reference photo
- Status — one of: requested → accepted → in progress → ready → delivered
- Price
- Payment status — e.g. deposit paid, balance paid
- Dates: when it was placed, when it's expected to be ready

---

## Open questions to settle before building
- Do measurement profiles need age brackets for girls, or exact measurements only?
- Can a tailor decline an order, and if so, does the customer get notified automatically or do they have to check?
- Is there a minimum deposit percentage, or does each tailor set their own?

---

*Next step: turn each of these into a Go `struct` — a Customer, Tailor, MeasurementProfile, and Order.*