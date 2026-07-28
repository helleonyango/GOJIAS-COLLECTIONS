## 1. Customer Registration Workflow

### Purpose

This workflow describes how a new customer creates an account and gains access to the GOJIAS COLLECTIONS platform.

### Workflow

Customer Visits Website
        │
        ▼
Clicks "Sign Up"
        │
        ▼
Enters Registration Details
(Name, Email, Phone, Password)
        │
        ▼
System Validates Input
        │
 ┌──────┴──────┐
 │             │
 ▼             ▼
Invalid      Valid
 │             │
 ▼             ▼
Display      Create
Errors       Account
 │             │
 └──────┬──────┘
        │
        ▼
Customer Receives Success Message
        │
        ▼
Customer Logs In
        │
        ▼
Customer Dashboard

### Business Rules

- Email address must be unique.
- Phone number must be unique.
- Password must meet minimum security requirements.
- Required fields cannot be left empty.
- Customer account is created only after successful validation.

## 2. Dressmaker Registration & Approval Workflow

### Purpose

This workflow describes how a dressmaker registers on the GOJIAS COLLECTIONS platform and how an administrator reviews and approves the application before the dressmaker can begin offering services.

### Workflow

Dressmaker Visits Website
        │
        ▼
Clicks "Register as a Dressmaker"
        │
        ▼
Enters Registration Details
(Personal Information, Business Information,
Portfolio, Location, Contact Details)
        │
        ▼
System Validates Input
        │
 ┌──────┴──────┐
 │             │
 ▼             ▼
Invalid      Valid
 │             │
 ▼             ▼
Display      Save Application
Errors       (Status = Pending)
                  │
                  ▼
      Administrator Reviews Application
                  │
          ┌───────┴────────┐
          │                │
          ▼                ▼
      Approved       Needs Changes/Rejected
          │                │
          ▼                ▼
Profile Becomes     Feedback Sent to
Visible to          Dressmaker
Customers               │
                         ▼
                Dressmaker Updates Information
                         │
                         ▼
                  Resubmits Application

                  ### Business Rules

- Email address must be unique.
- Phone number must be unique.
- Business name should be unique where possible.
- Required fields must be completed.
- At least one portfolio sample must be provided.
- Every new application starts with a **Pending** status.
- Only approved dressmakers are visible to customers.
- Dressmakers whose applications require changes may edit and resubmit their information.

## 3. Customer Places Order Workflow

### Purpose

This workflow describes how a customer browses the platform, selects a dressmaker or ready-made product, places an order, and submits it for processing.

### Workflow

Customer Logs In
        │
        ▼
Browse Dressmakers or Products
        │
        ▼
Select Dressmaker/Product
        │
        ▼
Choose Order Type
(Custom or Ready-made)
        │
        ▼
Provide Order Details
(Design Requirements or Product Selection)
        │
        ▼
Measurement Required?
        │
 ┌──────┴──────┐
 │             │
 ▼             ▼
Yes            No
 │             │
 ▼             ▼
Upload          Continue
Measurements
or Book
Appointment
        │
        ▼
Review Order Summary
        │
        ▼
Confirm Order
        │
        ▼
Order Created
(Status = Pending)
        │
        ▼
Dressmaker Receives Notification

## 4. Measurement Process Workflow

### Purpose

This workflow describes how measurements are collected for custom-made clothing to ensure garments fit the customer correctly.

### Workflow

Customer Places Custom Order
        │
        ▼
Measurement Required
        │
        ▼
Choose Measurement Method
        │
 ┌──────┼───────────────┐
 │      │               │
 ▼      ▼               ▼
Upload  Book        Request
Sheet   Appointment Tailor Visit
 │       │               │
 ▼       ▼               ▼
Measurements Collected
        │
        ▼
Dressmaker Reviews Measurements
        │
 ┌──────┴──────┐
 │             │
 ▼             ▼
Sufficient   Need Clarification
 │             │
 ▼             ▼
Approve      Request Updated
Measurements Measurements
 │             │
 └──────┬──────┘
        │
        ▼
Measurements Saved
        │
        ▼
Production Ready

### Business Rules

- Measurements are required for all custom-made orders.
- Customers may upload an existing measurement sheet.
- Customers may book an in-person measurement appointment.
- Customers may request a tailor visit for measurements.
- Dressmakers must review measurements before production begins.
- Production starts only after measurements are approved.

## 5. Appointment Booking Workflow

### Purpose

This workflow describes how customers schedule appointments with dressmakers for measurements, fittings, tailor visits, or garment pickup.

### Workflow

Customer Requests Appointment
        │
        ▼
Select Appointment Type
(Measurement, Fitting,
Tailor Visit, Pickup)
        │
        ▼
Choose Preferred Date & Time
        │
        ▼
System Checks Availability
        │
 ┌──────┴──────┐
 │             │
 ▼             ▼
Unavailable   Available
 │             │
 ▼             ▼
Suggest       Confirm
Alternative   Appointment
Date/Time
                  │
                  ▼
Appointment Saved
(Status = Scheduled)
                  │
                  ▼
Customer Receives Confirmation
                  │
                  ▼
Dressmaker Receives Notification
                  │
                  ▼
Appointment Day
                  │
                  ▼
Appointment Completed

### Business Rules

- Customer must be logged in to book an appointment.
- Appointment must be linked to an existing order.
- Appointment date and time must be available.
- Both customer and dressmaker receive confirmation after booking.
- Customers may reschedule or cancel appointments before the scheduled time.
- Appointment status is updated as it progresses (Scheduled, Confirmed, Completed, Cancelled, or Rescheduled).

## 6. Payment Process Workflow

### Purpose

This workflow describes how customers make payments for their orders and how the system verifies and records successful transactions.

### Workflow

Customer Confirms Order
        │
        ▼
Select Payment Method
(M-Pesa, Bank Transfer,
Cash on Delivery)
        │
        ▼
Submit Payment
        │
        ▼
System Verifies Payment
        │
 ┌──────┴──────┐
 │             │
 ▼             ▼
Failed      Successful
 │             │
 ▼             ▼
Display      Record Payment
Error        Update Payment Status
Message      (Successful)
 │             │
 └──────┬──────┘
        │
        ▼
Order Payment Status Updated
        │
        ▼
Customer Receives Confirmation
        │
        ▼
Dressmaker Receives Notification
        │
        ▼
Order Ready for Processing

### Business Rules

- Every payment must be linked to an existing order.
- Customers must choose one of the supported payment methods.
- Payment status must be recorded after verification.
- A successful payment updates the order payment status.
- Both the customer and the dressmaker receive payment confirmation.
- Cash on Delivery payments are marked as pending until payment is collected upon delivery or pickup.

- Ready-made products require full payment before shipment or pickup.
- Custom-made orders require a 50% deposit before production begins.
- The remaining 50% must be paid before the garment is collected or delivered.
## 7. Order Fulfillment Workflow

### Purpose

This workflow describes how a dressmaker processes a customer's order from acceptance to successful delivery or pickup.

### Workflow

Customer Places Order
        │
        ▼
Dressmaker Receives Order
        │
        ▼
Dressmaker Reviews Order
        │
 ┌──────┴──────┐
 │             │
 ▼             ▼
Reject      Accept
Order        Order
 │             │
 ▼             ▼
Notify      Confirm Order
Customer         │
                 ▼
      Measurements Complete?
                 │
        ┌────────┴────────┐
        │                 │
        ▼                 ▼
       No                Yes
        │                 │
        ▼                 ▼
Complete         Begin Production
Measurements            │
                        ▼
              Garment Production
                        │
                        ▼
               Quality Inspection
                        │
               ┌────────┴────────┐
               │                 │
               ▼                 ▼
          Needs Adjustment      Approved
               │                 │
               ▼                 ▼
      Make Adjustments     Ready for Pickup/
                           Delivery
                                │
                                ▼
                    Customer Pays Remaining Balance
                     (For Custom Orders)
                                │
                                ▼
                    Customer Collects Garment
                         or Receives Delivery
                                │
                                ▼
                      Order Marked Completed

        ### Business Rules

- Only approved dressmakers can process orders.
- Production begins only after required measurements are approved.
- Custom-made orders require a 50% deposit before production starts.
- Ready-made orders require full payment before shipment or pickup.
- Every completed order is marked as **Completed** in the system.
- Customers must pay the remaining balance before collecting or receiving a custom-made garment.              

## 8. Review Submission Workflow

### Purpose

This workflow describes how customers submit reviews and ratings after successfully receiving their completed orders.

### Workflow

Order Completed
        │
        ▼
Customer Receives Review Invitation
        │
        ▼
Customer Opens Review Page
        │
        ▼
Rate Dressmaker
(1–5 Stars)
        │
        ▼
Write Review
        │
        ▼
Submit Review
        │
        ▼
System Validates Review
        │
 ┌──────┴──────┐
 │             │
 ▼             ▼
Invalid      Valid
 │             │
 ▼             ▼
Display      Save Review
Errors           │
                 ▼
      Administrator Moderation
     (If Required)
                 │
         ┌───────┴────────┐
         │                │
         ▼                ▼
     Approved         Rejected
         │                │
         ▼                ▼
Review Published   Notify Customer
         │
         ▼
Dressmaker Rating Updated
### Business Rules

- Only customers with completed orders can submit reviews.
- Each completed order can receive only one review.
- Ratings must be between 1 and 5 stars.
- Reviews must comply with platform guidelines.
- Administrators may moderate reported or inappropriate reviews.
- Published reviews contribute to the dressmaker's overall rating.

## 9. Administrator Approval Process Workflow

### Purpose

This workflow describes how administrators review dressmaker applications, moderate platform content, resolve issues, and maintain the quality and integrity of GOJIAS COLLECTIONS.

### Workflow

New Application or Report Received
                │
                ▼
Administrator Logs In
                │
                ▼
Review Submitted Information
                │
                ▼
Decision Required
                │
        ┌───────┴────────┐
        │                │
        ▼                ▼
     Approve      Reject / Request Changes
        │                │
        ▼                ▼
Update Status      Send Feedback
        │                │
        └───────┬────────┘
                │
                ▼
System Updates Records
                │
                ▼
Notify Customer or Dressmaker
                │
                ▼
Process Completed

### Business Rules

- Only administrators can approve or reject dressmaker applications.
- Every decision must be recorded in the system.
- Administrators may request additional information before making a decision.
- Administrators may moderate reviews that violate platform guidelines.
- Users must be notified whenever the status of their application or review changes.
- All administrative actions should be logged for accountability.