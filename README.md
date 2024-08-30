# Service Account API

## Overview

The Service Account API is a Go-based application designed to handle account-related operations such as registration, deposits, withdrawals, and balance inquiries. It uses the Echo framework for routing and GORM for ORM with a PostgreSQL database.

## Features

- Register new accounts
- Deposit funds into accounts
- Withdraw funds from accounts
- Retrieve account balances

## Prerequisites

- Go 1.19 or higher
- Docker (for containerized deployment)
- PostgreSQL (for database)

## Getting Started

### Clone the Repository

First, clone the repository to your local machine:

```bash
git clone https://github.com/runydek/service-account.git
cd service-account
