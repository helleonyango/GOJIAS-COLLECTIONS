# GOJIAS COLLECTIONS Database Design
## 1. Database Overview

The GOJIAS COLLECTIONS database is designed to store and manage all information required for customers, dressmakers, administrators, products, orders, appointments, payments, reviews, and notifications.

The database aims to ensure data integrity, security, scalability, and efficient retrieval of information while supporting the platform's business processes.

## 2. Main Database Entities
The Version 1 database consists of the following entities:

1. Customers
2. Dressmakers
3. Administrators
4. Products
5. Orders
6. Measurements
7. Appointments
8. Payments
9. Reviews
10. Notifications

## 3. Customers Table

### Purpose

Stores customer account information and profile details.

### Fields

| Field Name | Data Type | Description |
|------------|-----------|-------------|
| customer_id | UUID | Unique identifier for each customer |
| first_name | VARCHAR(100) | Customer's first name |
| last_name | VARCHAR(100) | Customer's last name |
| email | VARCHAR(255) | Customer's email address |
| phone_number | VARCHAR(20) | Customer's phone number |
| password_hash | VARCHAR(255) | Encrypted password |
| profile_photo | VARCHAR(255) | Profile image URL or file path |
| country | VARCHAR(100) | Customer's country |
| city | VARCHAR(100) | Customer's city |
| address | TEXT | Delivery address |
| created_at | TIMESTAMP | Account creation date |
| updated_at | TIMESTAMP | Last profile update |

### Primary Key

- customer_id

### Foreign Keys

- None

### Relationships
- One customer can place many orders.
- One customer can make many payments.
- One customer can write many reviews.
- One customer can upload multiple measurement records.
- One customer can book multiple appointments.

## 4. Dressmakers Table

### Purpose

Stores dressmaker account information, business details, verification status, and professional portfolio.

### Fields

| Field Name | Data Type | Description |
|------------|-----------|-------------|
| dressmaker_id | UUID | Unique identifier for each dressmaker |
| first_name | VARCHAR(100) | Dressmaker's first name |
| last_name | VARCHAR(100) | Dressmaker's last name |
| business_name | VARCHAR(150) | Registered business or brand name |
| email | VARCHAR(255) | Dressmaker's email address |
| phone_number | VARCHAR(20) | Contact phone number |
| password_hash | VARCHAR(255) | Encrypted password |
| profile_photo | VARCHAR(255) | Profile image URL or file path |
| portfolio | TEXT | Links or references to portfolio images |
| years_of_experience | INTEGER | Number of years of tailoring experience |
| specialization | VARCHAR(150) | Area of specialization (e.g., Ankara, Bridal Wear, Kitenge, Casual Wear) |
| country | VARCHAR(100) | Country of operation |
| city | VARCHAR(100) | City of operation |
| address | TEXT | Business address |
| verification_status | VARCHAR(20) | Pending, Approved, Rejected |
| bio | TEXT | Brief introduction about the dressmaker |
| average_rating | DECIMAL(2,1) | Average customer rating |
| total_reviews | INTEGER | Total number of customer reviews |
| created_at | TIMESTAMP | Account creation date |
| updated_at | TIMESTAMP | Last profile update |

### Primary Key

- dressmaker_id

### Foreign Keys

- None

### Relationships

- One dressmaker can receive many customer orders.
- One dressmaker can create many product listings.
- One dressmaker can manage many appointments.
- One dressmaker can receive many customer reviews.
- One dressmaker can receive many payments for completed orders.

## 5. Administrators Table

### Purpose

Stores administrator account information, roles, and access permissions for managing the platform.

### Fields

| Field Name | Data Type | Description |
|------------|-----------|-------------|
| admin_id | UUID | Unique identifier for each administrator |
| first_name | VARCHAR(100) | Administrator's first name |
| last_name | VARCHAR(100) | Administrator's last name |
| email | VARCHAR(255) | Administrator's email address |
| phone_number | VARCHAR(20) | Administrator's contact number |
| password_hash | VARCHAR(255) | Encrypted password |
| role | VARCHAR(50) | Administrator role (e.g., Super Admin, Support Admin, Content Admin) |
| profile_photo | VARCHAR(255) | Profile image URL or file path |
| account_status | VARCHAR(20) | Active or Inactive |
| last_login | TIMESTAMP | Date and time of the last successful login |
| created_at | TIMESTAMP | Account creation date |
| updated_at | TIMESTAMP | Last account update |

### Primary Key

- admin_id

### Foreign Keys

- None

### Relationships

- One administrator can review many dressmaker applications.
- One administrator can manage many customer accounts.
- One administrator can manage many dressmaker accounts.
- One administrator can moderate many reviews.
- One administrator can publish many announcements.

yes product table

## 6. Products Table

### Purpose

Stores information about ready-made clothing listed by dressmakers for customers to browse and purchase.

### Fields

| Field Name | Data Type | Description |
|------------|-----------|-------------|
| product_id | UUID | Unique identifier for each product |
| dressmaker_id | UUID | References the dressmaker who owns the product |
| product_name | VARCHAR(150) | Name of the product |
| category | VARCHAR(100) | Product category (e.g., Dresses, Skirts, Tops, Suits) |
| fabric_type | VARCHAR(100) | Fabric used (e.g., Ankara, Kitenge, Kitenge Mix, Cotton) |
| description | TEXT | Detailed product description |
| price | DECIMAL(10,2) | Selling price |
| stock_quantity | INTEGER | Number of items available |
| available_sizes | VARCHAR(100) | Available sizes (e.g., S, M, L, XL) |
| primary_image | VARCHAR(255) | Main product image URL or file path |
| product_status | VARCHAR(20) | Available, Out of Stock, Hidden |
| created_at | TIMESTAMP | Date the product was added |
| updated_at | TIMESTAMP | Date the product was last updated |

### Primary Key

- product_id

### Foreign Keys

- dressmaker_id → Dressmakers(dressmaker_id)

### Relationships

- One dressmaker can list many products.
- One product belongs to one dressmaker.
- One product can appear in many customer orders.

## 7. Orders Table

### Purpose

Stores information about customer orders, including custom tailoring requests and ready-made clothing purchases.

### Fields

| Field Name | Data Type | Description |
|------------|-----------|-------------|
| order_id | UUID | Unique identifier for each order |
| customer_id | UUID | References the customer who placed the order |
| dressmaker_id | UUID | References the dressmaker fulfilling the order |
| product_id | UUID | References the ready-made product (nullable for custom orders) |
| order_type | VARCHAR(20) | Custom or Ready-made |
| order_description | TEXT | Customer's design requirements or additional notes |
| total_amount | DECIMAL(10,2) | Total amount payable |
| payment_status | VARCHAR(20) | Pending, Paid, Partially Paid, Refunded |
| order_status | VARCHAR(30) | Pending, Accepted, In Progress, Ready for Fitting, Ready for Pickup, Completed, Cancelled |
| delivery_method | VARCHAR(30) | Pickup, Delivery |
| expected_completion_date | DATE | Estimated completion date |
| completed_at | TIMESTAMP | Date and time the order was completed |
| created_at | TIMESTAMP | Date the order was placed |
| updated_at | TIMESTAMP | Last update to the order |

### Primary Key

- order_id

### Foreign Keys

- customer_id → Customers(customer_id)
- dressmaker_id → Dressmakers(dressmaker_id)
- product_id → Products(product_id)

### Relationships

- One customer can place many orders.
- One dressmaker can receive many orders.
- One product can appear in many orders.
- One order can have one payment record.
- One order can have one measurement record (for custom orders).
- One order can have one or more appointments.
- One completed order can receive one customer review.

## 8. Measurements Table

### Purpose

Stores customer body measurements used by dressmakers to produce well-fitting custom-made garments.

### Fields

| Field Name | Data Type | Description |
|------------|-----------|-------------|
| measurement_id | UUID | Unique identifier for each measurement record |
| customer_id | UUID | References the customer |
| order_id | UUID | References the related custom order |
| measurement_type | VARCHAR(20) | Uploaded Sheet, In-Person, Tailor Visit |
| height | DECIMAL(5,2) | Customer's height (cm) |
| bust | DECIMAL(5,2) | Bust measurement (cm) |
| waist | DECIMAL(5,2) | Waist measurement (cm) |
| hips | DECIMAL(5,2) | Hip measurement (cm) |
| shoulder | DECIMAL(5,2) | Shoulder width (cm) |
| sleeve_length | DECIMAL(5,2) | Sleeve length (cm) |
| dress_length | DECIMAL(5,2) | Dress or garment length (cm) |
| inseam | DECIMAL(5,2) | Inseam measurement (cm), where applicable |
| neck | DECIMAL(5,2) | Neck circumference (cm) |
| additional_notes | TEXT | Extra fitting instructions or special requirements |
| measurement_file | VARCHAR(255) | Uploaded measurement sheet file path or URL |
| measured_by | VARCHAR(30) | Customer, Dressmaker, or Tailor |
| measured_at | TIMESTAMP | Date measurements were taken |
| created_at | TIMESTAMP | Date the record was created |
| updated_at | TIMESTAMP | Last update date |

### Primary Key

- measurement_id

### Foreign Keys

- customer_id → Customers(customer_id)
- order_id → Orders(order_id)

### Relationships

- One customer can have multiple measurement records.
- One order can have one associated measurement record.
- One measurement record belongs to one customer.

## 9. Appointments Table

### Purpose

Stores appointment information for customer measurements, fittings, tailor visits, and garment collection.

### Fields

| Field Name | Data Type | Description |
|------------|-----------|-------------|
| appointment_id | UUID | Unique identifier for each appointment |
| order_id | UUID | References the related order |
| customer_id | UUID | References the customer |
| dressmaker_id | UUID | References the dressmaker |
| appointment_type | VARCHAR(30) | Measurement, Fitting, Tailor Visit, Pickup |
| appointment_date | DATE | Scheduled appointment date |
| appointment_time | TIME | Scheduled appointment time |
| location | TEXT | Appointment location |
| status | VARCHAR(20) | Scheduled, Confirmed, Completed, Cancelled, Rescheduled |
| notes | TEXT | Additional appointment notes |
| created_at | TIMESTAMP | Date the appointment was created |
| updated_at | TIMESTAMP | Date the appointment was last updated |

### Primary Key

- appointment_id

### Foreign Keys

- order_id → Orders(order_id)
- customer_id → Customers(customer_id)
- dressmaker_id → Dressmakers(dressmaker_id)

### Relationships

- One customer can have many appointments.
- One dressmaker can have many appointments.
- One ord## 10. Payments Table

### Purpose

Stores payment information for customer orders, including payment methods, transaction status, and payment history.

### Fields

| Field Name | Data Type | Description |
|------------|-----------|-------------|
| payment_id | UUID | Unique identifier for each payment |
| order_id | UUID | References the related order |
| customer_id | UUID | References the customer making the payment |
| dressmaker_id | UUID | References the dressmaker receiving the payment |
| payment_method | VARCHAR(30) | M-Pesa, Cash on Delivery, Bank Transfer |
| transaction_reference | VARCHAR(100) | Payment transaction reference number |
| amount | DECIMAL(10,2) | Amount paid |
| currency | VARCHAR(10) | Currency used (e.g., KES, USD) |
| payment_status | VARCHAR(20) | Pending, Successful, Failed, Refunded |
| payment_date | TIMESTAMP | Date and time the payment was made |
| notes | TEXT | Additional payment notes |
| created_at | TIMESTAMP | Date the payment record was created |
| updated_at | TIMESTAMP | Date the payment record was last updated |

### Primary Key

- payment_id

### Foreign Keys

- order_id → Orders(order_id)
- customer_id → Customers(customer_id)
- dressmaker_id → Dressmakers(dressmaker_id)

### Relationships

- One customer can make many payments.
- One dressmaker can receive many payments.
- One order can have one or more payment records.er can have one or more appointments.

## 10. Payments Table

### Purpose

Stores payment information for customer orders, including payment methods, transaction status, and payment history.

### Fields

| Field Name | Data Type | Description |
|------------|-----------|-------------|
| payment_id | UUID | Unique identifier for each payment |
| order_id | UUID | References the related order |
| customer_id | UUID | References the customer making the payment |
| dressmaker_id | UUID | References the dressmaker receiving the payment |
| payment_method | VARCHAR(30) | M-Pesa, Cash on Delivery, Bank Transfer |
| transaction_reference | VARCHAR(100) | Payment transaction reference number |
| amount | DECIMAL(10,2) | Amount paid |
| currency | VARCHAR(10) | Currency used (e.g., KES, USD) |
| payment_status | VARCHAR(20) | Pending, Successful, Failed, Refunded |
| payment_date | TIMESTAMP | Date and time the payment was made |
| notes | TEXT | Additional payment notes |
| created_at | TIMESTAMP | Date the payment record was created |
| updated_at | TIMESTAMP | Date the payment record was last updated |

### Primary Key

- payment_id

### Foreign Keys

- order_id → Orders(order_id)
- customer_id → Customers(customer_id)
- dressmaker_id → Dressmakers(dressmaker_id)

### Relationships

- One customer can make many payments.
- One dressmaker can receive many payments.
- One order can have one or more payment records.

## 11. Reviews Table

### Purpose

Stores customer ratings and reviews for completed orders, helping maintain quality standards and build trust within the platform.

### Fields

| Field Name | Data Type | Description |
|------------|-----------|-------------|
| review_id | UUID | Unique identifier for each review |
| order_id | UUID | References the completed order |
| customer_id | UUID | References the customer who wrote the review |
| dressmaker_id | UUID | References the dressmaker being reviewed |
| rating | INTEGER | Rating score (1–5 stars) |
| review_title | VARCHAR(150) | Short title for the review |
| review_comment | TEXT | Customer's written feedback |
| review_status | VARCHAR(20) | Pending, Approved, Rejected |
| admin_id | UUID | References the administrator who reviewed or moderated the review (nullable) |
| created_at | TIMESTAMP | Date the review was submitted |
| updated_at | TIMESTAMP | Date the review was last updated |

### Primary Key

- review_id

### Foreign Keys

- order_id → Orders(order_id)
- customer_id → Customers(customer_id)
- dressmaker_id → Dressmakers(dressmaker_id)
- admin_id → Administrators(admin_id)

### Relationships

- One customer can write many reviews.
- One dressmaker can receive many reviews.
- One completed order can have one review.
- One administrator can moderate many reviews.
Notification

## 12. Notifications Table

### Purpose

Stores notifications sent to customers, dressmakers, and administrators regarding important platform activities and updates.

### Fields

| Field Name | Data Type | Description |
|------------|-----------|-------------|
| notification_id | UUID | Unique identifier for each notification |
| recipient_type | VARCHAR(20) | Customer, Dressmaker, or Administrator |
| customer_id | UUID | References the customer (nullable) |
| dressmaker_id | UUID | References the dressmaker (nullable) |
| admin_id | UUID | References the administrator (nullable) |
| title | VARCHAR(150) | Notification title |
| message | TEXT | Notification content |
| notification_type | VARCHAR(30) | Order Update, Payment, Appointment, Review, System |
| is_read | BOOLEAN | Indicates whether the notification has been read |
| sent_at | TIMESTAMP | Date and time the notification was sent |
| created_at | TIMESTAMP | Date the notification record was created |

### Primary Key

- notification_id

### Foreign Keys

- customer_id → Customers(customer_id)
- dressmaker_id → Dressmakers(dressmaker_id)
- admin_id → Administrators(admin_id)

### Relationships

- One customer can receive many notifications.
- One dressmaker can receive many notifications.
- One administrator can receive many notifications.

## 13. Entity Relationships

The GOJIAS COLLECTIONS database is built around the **Orders** table, which connects customers, dressmakers, products, measurements, appointments, payments, and reviews.

### Relationships

- One Customer can place many Orders.
- One Dressmaker can fulfill many Orders.
- One Dressmaker can list many Products.
- One Product belongs to one Dressmaker.
- One Customer can have many Measurement records.
- One Order can have one Measurement record.
- One Order can have one or more Appointments.
- One Customer can make many Payments.
- One Order can have one or more Payment records.
- One Customer can write many Reviews.
- One Dressmaker can receive many Reviews.
- One Completed Order can have one Review.
- One Customer can receive many Notifications.
- One Dressmaker can receive many Notifications.
- One Administrator can receive many Notifications.

# GOJIAS COLLECTIONS Workflow Diagrams
