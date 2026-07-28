# Product Requirements Document (PRD)

## Project Name
GOJIAS COLLECTIONS

## Version
1.0

## Status
Planning

## Prepared By
Hellen

## Date
11 July 2026## 1.

## 1. Product Overview
GOJIAS COLLECTIONS is a fashion technology platform that connects customers with skilled dressmakers while offering both ready-made and custom-made African-inspired clothing. The platform enables customers to browse designs, discover dressmakers by location, upload measurements, book fitting appointments, and order garments conveniently online.

The platform also provides dressmakers with a digital space to showcase their work, receive customer orders, and grow their businesses while promoting local craftsmanship.

## 2. Problem Statement
Many customers struggle to find skilled and trustworthy dressmakers who match their style, location, and budget. Communication is often done through phone calls or messaging apps, making it difficult to share measurements, track orders, and monitor progress. Customers also have limited access to ready-made African-inspired clothing from local artisans.

On the other hand, many talented dressmakers have limited opportunities to market their work beyond their local communities. They often rely on word-of-mouth referrals, making it difficult to reach new customers and grow their businesses.

GOJIAS COLLECTIONS addresses these challenges by providing a single digital platform where customers can discover dressmakers, browse ready-made clothing, place custom orders, upload measurements, schedule fitting appointments, and support local artisans through a convenient and trusted online experience.

## 3. Product Objectives

The primary objectives of GOJIAS COLLECTIONS are to:

- Provide customers with a convenient platform for ordering both ready-made and custom-made African-inspired clothing.
- Connect customers with skilled and trusted dressmakers based on their location.
- Enable customers to upload measurement sheets and schedule fitting appointments.
- Help dressmakers showcase their work, receive orders, and expand their customer base.
- Promote local artisans by increasing their online visibility and business opportunities.
- Deliver an affordable, reliable, and user-friendly shopping experience.
- Build a scalable platform that can expand to multiple countries over time.

## 4. Target Users
GOJIAS COLLECTIONS serves three main groups of users:

### Customers
Women and girls looking for ready-made or custom-made African-inspired clothing. Customers can browse designs, search for dressmakers by location, upload measurements, book fitting appointments, and place orders.

### Dressmakers
Professional tailors and fashion designers who create and sell custom-made or ready-made clothing. They can showcase their portfolios, receive customer orders, manage appointments, and grow their businesses through the platform.

### Administrators
Platform administrators responsible for managing users, monitoring orders, verifying dressmakers, handling customer support, and maintaining the overall quality and security of the platform.

## 5. Project Scope

### In Scope (Version 1)

The first version of GOJIAS COLLECTIONS will include:

- Customer registration and login
- Dressmaker registration and profile management
- Browse ready-made clothing
- Search dressmakers by location
- View dressmaker portfolios
- Place custom clothing orders
- Upload measurement sheets
- Book fitting appointments
- Order management and status tracking
- Payment options:
  - M-Pesa
  - Cash on Delivery
  - Bank Transfer
- Administrator dashboard for managing users and orders

### Out of Scope (Future Versions)

The following features will be considered for future releases:

- Dedicated mobile application
- Live chat between customers and dressmakers
- Delivery tracking
- AI-powered fashion recommendations
- Multi-language support
- Fabric marketplace
- International shipping

## 6. Functional Requirements
### 6.1 Customer Functional Requirements

The system shall provide the following functionality for customers:

**FR-C-01: Customer Registration**
- The system shall allow customers to create an account using their email address and password.

**FR-C-02: Customer Login**
- The system shall allow registered customers to log in securely.

**FR-C-03: Password Recovery**
- The system shall allow customers to reset forgotten passwords.

**FR-C-04: Profile Management**
- The system shall allow customers to view and update their personal profile information.

**FR-C-05: Search Dressmakers**
- The system shall allow customers to search for dressmakers based on location.

**FR-C-06: View Dressmaker Profiles**
- The system shall allow customers to view dressmaker profiles, portfolios, ratings, and reviews.

**FR-C-07: Browse Ready-Made Clothing**
- The system shall allow customers to browse available ready-made clothing.

**FR-C-08: View Product Details**
- The system shall display product descriptions, prices, available sizes, and images.

**FR-C-09: Place Custom Orders**
- The system shall allow customers to place custom clothing orders.

**FR-C-10: Upload Measurements**
- The system shall allow customers to upload measurement sheets during the ordering process.

**FR-C-11: Book Fitting Appointments**
- The system shall allow customers to schedule fitting appointments with dressmakers.

**FR-C-12: Make Payments**
- The system shall support payments using:
  - M-Pesa
  - Cash on Delivery
  - Bank Transfer

**FR-C-13: Track Orders**
- The system shall allow customers to track the status of their orders.

**FR-C-14: Order History**
- The system shall allow customers to view their previous and current orders.

**FR-C-15: Notifications**
- The system shall notify customers of important order updates.

**FR-C-16: Ratings and Reviews**
- The system shall allow customers to rate and review completed orders and dressmakers.

**FR-C-17: Save Favorites**
- The system shall allow customers to save favorite dressmakers and clothing items for future reference.

### 6.2 Dressmaker Functional Requirements

The system shall provide the following functionality for dressmakers:

**FR-D-01: Dressmaker Registration**
- The system shall allow dressmakers to create an account.

**FR-D-02: Verification Application**
- The system shall require dressmakers to submit their business information, location, portfolio, contact details, and experience for administrator review.

**FR-D-03: Administrator Approval**
- The system shall ensure that only administrator-approved dressmakers can publish their profiles and receive customer orders.

**FR-D-04: Profile Management**
- The system shall allow dressmakers to create, view, and update their professional profiles.

**FR-D-05: Portfolio Management**
- The system shall allow dressmakers to upload, edit, and remove images of their completed work.

**FR-D-06: Product Management**
- The system shall allow dressmakers to add, edit, update, and remove ready-made clothing listings.

**FR-D-07: Order Management**
- The system shall allow dressmakers to receive, accept, decline, and manage customer orders.

**FR-D-08: Appointment Management**
- The system shall allow dressmakers to view, accept, reschedule, or decline fitting appointments.

**FR-D-09: Order Status Updates**
- The system shall allow dressmakers to update the progress of customer orders (e.g., Pending, In Progress, Ready for Fitting, Completed).

**FR-D-10: Customer Measurements**
- The system shall allow dressmakers to securely view customer measurement sheets associated with their orders.

**FR-D-11: Payment Records**
- The system shall allow dressmakers to view payment information and payment status for their orders.

**FR-D-12: Notifications**
- The system shall notify dressmakers of new orders, appointment requests, payments, and customer messages.

**FR-D-13: Ratings and Reviews**
- The system shall allow dressmakers to view customer ratings and reviews received after completed orders.

**FR-D-14: Dashboard**
- The system shall provide dressmakers with a dashboard displaying active orders, completed orders, appointments, earnings, and customer reviews.

**FR-D-15: Availability Management**
- The system shall allow dressmakers to set their availability for accepting new orders and booking fitting appointments.

### 6.3 Administrator Functional Requirements

The system shall provide the following functionality for administrators:

**FR-A-01: Administrator Login**
- The system shall allow administrators to log in securely.

**FR-A-02: Dashboard**
- The system shall provide administrators with a dashboard displaying platform statistics, recent activities, pending approvals, orders, and user summaries.

**FR-A-03: Dressmaker Verification**
- The system shall allow administrators to review dressmaker applications and approve, request changes, or reject applications.

**FR-A-04: User Management**
- The system shall allow administrators to view, manage, suspend, or deactivate customer and dressmaker accounts when necessary.

**FR-A-05: Order Management**
- The system shall allow administrators to monitor all customer orders and intervene when necessary to resolve disputes or issues.

**FR-A-06: Product Management**
- The system shall allow administrators to review, edit, or remove inappropriate product listings.

**FR-A-07: Appointment Monitoring**
- The system shall allow administrators to monitor scheduled fitting appointments and assist in resolving appointment conflicts.

**FR-A-08: Payment Monitoring**
- The system shall allow administrators to monitor payment transactions and payment statuses.

**FR-A-09: Review Moderation**
- The system shall allow administrators to review, approve, or remove inappropriate customer reviews and ratings.

**FR-A-10: Notifications**
- The system shall notify administrators of important platform events, including new dressmaker applications, reported issues, failed payments, and customer complaints.

**FR-A-11: Reports and Analytics**
- The system shall generate reports on users, orders, sales, payments, and platform performance.

**FR-A-12: Content Management**
- The system shall allow administrators to manage homepage content, promotional banners, featured dressmakers, and featured products.

**FR-A-13: Customer Support**
- The system shall allow administrators to receive, track, and respond to customer and dressmaker support requests.

**FR-A-14: Platform Configuration**
- The system shall allow administrators to manage platform settings such as payment methods, service areas, categories, and system preferences.

**FR-D-15: Availability Management**
- The system shall allow dressmakers to set their availability for accepting new orders and booking fitting appointments.

## 7. Non-Functional Requirements

The following non-functional requirements define the quality, performance, security, and reliability expectations for GOJIAS COLLECTIONS.

### NFR-01: Performance
- The platform shall load pages within 3 seconds under normal operating conditions.
- Search results shall be displayed within 2 seconds.

### NFR-02: Availability
- The platform shall be available 24 hours a day, 7 days a week, except during scheduled maintenance.

### NFR-03: Security
- User passwords shall be securely encrypted before being stored.
- The platform shall use secure authentication and authorization mechanisms.
- Customer personal information and measurement data shall be protected against unauthorized access.

### NFR-04: Usability
- The platform shall provide an intuitive and user-friendly interface.
- Users shall be able to complete common tasks with minimal training.
- The website shall be responsive and usable on desktops, tablets, and mobile devices.

### NFR-05: Reliability
- The platform shall maintain accurate customer, dressmaker, order, and payment records.
- Data shall not be lost during normal operation.

### NFR-06: Scalability
- The platform shall support future expansion to additional countries and increasing numbers of users without major redesign.

### NFR-07: Maintainability
- The system shall be modular and well documented to simplify future maintenance and feature development.

### NFR-08: Compatibility
- The platform shall support modern web browsers including Google Chrome, Mozilla Firefox, Microsoft Edge, and Safari.

### NFR-09: Backup and Recovery
- The platform shall perform regular database backups.
- The system shall support data recovery in the event of accidental data loss or system failure.

### NFR-10: Privacy
- Customer, dressmaker, and payment information shall be handled in accordance with applicable data protection and privacy regulations.

### NFR-11: Accessibility
- The platform shall follow accessibility best practices to ensure that people with different abilities can use the system effectively.

## 8. User Stories
### 8.1 Customer User Stories

**US-C-01**
As a customer, I want to create an account so that I can place orders and manage my profile.

**US-C-02**
As a customer, I want to log in securely so that my personal information is protected.

**US-C-03**
As a customer, I want to search for dressmakers by location so that I can find someone near me.

**US-C-04**
As a customer, I want to browse dressmaker portfolios so that I can choose a dressmaker whose style I like.

**US-C-05**
As a customer, I want to browse ready-made clothing so that I can purchase outfits immediately if I do not need custom tailoring.

**US-C-06**
As a customer, I want to upload my measurement sheet so that my clothes fit me correctly.

**US-C-07**
As a customer, I want to book a fitting appointment so that my measurements can be verified when necessary.

**US-C-08**
As a customer, I want to place a custom clothing order so that I can receive clothing made to my preferences.

**US-C-09**
As a customer, I want to pay using my preferred payment method so that the checkout process is convenient.

**US-C-10**
As a customer, I want to track my order so that I know its current status.

**US-C-11**
As a customer, I want to receive notifications so that I stay informed about my order progress.

**US-C-12**
As a customer, I want to rate and review my dressmaker so that I can share my experience and help other customers make informed decisions.

### 8.2 Dressmaker User Stories

**US-D-01**
As a dressmaker, I want to create an account so that I can join the GOJIAS COLLECTIONS platform.

**US-D-02**
As a dressmaker, I want to submit my business information and portfolio for verification so that customers can trust my profile.

**US-D-03**
As a dressmaker, I want to receive approval from an administrator so that I can begin offering my services on the platform.

**US-D-04**
As a dressmaker, I want to manage my profile so that customers always see my latest information and services.

**US-D-05**
As a dressmaker, I want to upload and update my portfolio so that I can showcase my skills and attract more customers.

**US-D-06**
As a dressmaker, I want to list my ready-made clothing so that customers can purchase my products online.

**US-D-07**
As a dressmaker, I want to receive notifications when customers place orders so that I can respond promptly.

**US-D-08**
As a dressmaker, I want to accept or decline customer orders so that I can manage my workload effectively.

**US-D-09**
As a dressmaker, I want to schedule and manage fitting appointments so that I can provide accurate measurements and quality service.

**US-D-10**
As a dressmaker, I want to update the status of customer orders so that customers know the progress of their garments.

**US-D-11**
As a dressmaker, I want to view customer measurement sheets so that I can create garments that fit correctly.

**US-D-12**
As a dressmaker, I want to view payment information so that I can confirm payment before processing orders when required.

**US-D-13**
As a dressmaker, I want to view customer ratings and reviews so that I can improve my services and build my reputation.

**US-D-14**
As a dressmaker, I want to manage my availability so that customers can only book appointments and place orders when I am available.

**US-D-15**
As a dressmaker, I want to access a dashboard so that I can monitor my orders, appointments, earnings, and overall business performance.
### 8.3 Administrator User Stories

**US-A-01**
As an administrator, I want to log in securely so that I can manage the platform safely.

**US-A-02**
As an administrator, I want to review dressmaker applications so that only qualified and verified dressmakers are visible to customers.

**US-A-03**
As an administrator, I want to approve, request changes, or reject dressmaker applications so that the platform maintains high service quality.

**US-A-04**
As an administrator, I want to manage customer and dressmaker accounts so that I can handle policy violations and maintain platform integrity.

**US-A-05**
As an administrator, I want to monitor customer orders so that I can assist in resolving disputes when necessary.

**US-A-06**
As an administrator, I want to monitor payment transactions so that I can ensure payments are processed correctly.

**US-A-07**
As an administrator, I want to manage product listings so that inappropriate or inaccurate content can be removed.

**US-A-08**
As an administrator, I want to moderate customer reviews and ratings so that the platform remains respectful and trustworthy.

**US-A-09**
As an administrator, I want to respond to customer and dressmaker support requests so that users receive timely assistance.

**US-A-10**
As an administrator, I want to publish announcements and promotional messages so that I can communicate important information to users.

**US-A-11**
As an administrator, I want to view reports and analytics so that I can evaluate the platform's performance and make informed decisions.

**US-A-12**
As an administrator, I want to configure platform settings so that I can manage payment methods, service areas, and other operational preferences.

## 9. Success Metrics

The success of GOJIAS COLLECTIONS will be measured using the following key performance indicators (KPIs):

### User Growth
- Register at least 1,000 customers within the first year.
- Register at least 100 verified dressmakers within the first year.

### Customer Engagement
- Achieve a high percentage of returning customers.
- Increase the average number of orders placed per customer over time.

### Order Performance
- Maintain a high order completion rate.
- Minimize order cancellations and disputes.

### Customer Satisfaction
- Achieve an average customer rating of 4.5 stars or higher.
- Receive positive customer feedback regarding quality, convenience, and reliability.

### Business Growth
- Increase the number of ready-made clothing listings available on the platform.
- Expand the platform to additional regions and countries over time.

### Platform Reliability
- Maintain high platform availability with minimal downtime.
- Resolve customer support requests within an acceptable timeframe.

Register at least 500 customers within the first year.

Register at least 50 verified dressmakers within the first year.

## 10. Future Enhancements

Future versions of GOJIAS COLLECTIONS may include the following enhancements:

- Dedicated Android and iOS mobile applications.
- Online card payment integration (Visa, Mastercard, etc.).
- International payment support for global customers.
- Live chat between customers and dressmakers.
- Real-time order and delivery tracking.
- AI-powered clothing and dressmaker recommendations based on customer preferences.
- Multi-language support to serve users from different regions.
- Fabric marketplace where customers can purchase fabrics from trusted suppliers.
- Virtual measurement assistance using image processing or body scanning technologies.
- Loyalty and rewards program for returning customers.
- Discount coupons and promotional campaigns.
- Referral program for customers and dressmakers.
- Integration with courier and delivery services.
- Advanced analytics and business insights for dressmakers.
- Expansion to additional countries and regions.


