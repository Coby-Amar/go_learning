-- +goose Up
CREATE EXTENSION "uuid-ossp";

CREATE TABLE _users (
    _id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(), 
    _created_at TIMESTAMP NOT NULL DEFAULT NOW(), 
    _updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    _last_login TIMESTAMP NOT NULL DEFAULT NOW(),
    _name TEXT NOT NULL CHECK(CHAR_LENGTH(_name) > 0),
    _email TEXT NOT NULL UNIQUE CHECK(CHAR_LENGTH(_name) > 0),
    _phone_number TEXT NOT NULL CHECK(CHAR_LENGTH(_name) > 0)
);

CREATE TABLE _vault (
    _user_id UUID NOT NULL PRIMARY KEY, 
    _hashed_pw CHAR(64) NOT NULL,
    _active BOOLEAN NOT NULL DEFAULT TRUE,
    FOREIGN KEY (_user_id) REFERENCES _users(_id) ON DELETE CASCADE ON UPDATE NO ACTION
);

CREATE TABLE _daily_limit (
    _user_id UUID NOT NULL PRIMARY KEY,
    _carbohydrate SMALLINT NOT NULL CHECK(_carbohydrate > 0), 
    _protein SMALLINT NOT NULL CHECK(_protein > 0),
    _fat SMALLINT NOT NULL CHECK(_fat > 0),
    FOREIGN KEY (_user_id) REFERENCES _users(_id) ON DELETE CASCADE ON UPDATE NO ACTION
);

CREATE TABLE _products (
    _id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(), 
    _created_at TIMESTAMP NOT NULL DEFAULT NOW(), 
    _updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    _name TEXT NOT NULL UNIQUE CHECK(CHAR_LENGTH(_name) > 0),
    _amount SMALLINT NOT NULL CHECK(_amount > 0),
    _carbohydrate SMALLINT NOT NULL CHECK(_carbohydrate > -1), 
    _protein SMALLINT NOT NULL CHECK(_protein > -1),
    _fat SMALLINT NOT NULL CHECK(_fat > -1),
    _user_id UUID NOT NULL,
    FOREIGN KEY (_user_id) REFERENCES _users(_id) ON DELETE CASCADE ON UPDATE NO ACTION
);

CREATE INDEX _products_fkey_user_id ON _products(_user_id);

CREATE TABLE _reports (
    _id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    _created_at TIMESTAMP NOT NULL DEFAULT NOW(), 
    _updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    _date DATE NOT NULL UNIQUE,
    _amout_of_entries SMALLINT NOT NULL CHECK(_amout_of_entries > 0),
    _carbohydrates_total NUMERIC(4,2) NOT NULL DEFAULT 0.00 CHECK(_carbohydrates_total > -1), 
    _proteins_total NUMERIC(4,2) NOT NULL DEFAULT 0.00 CHECK(_proteins_total > -1),
    _fats_total NUMERIC(4,2) NOT NULL DEFAULT 0.00 CHECK(_fats_total > -1),
    _user_id UUID NOT NULL,
    FOREIGN KEY (_user_id) REFERENCES _users(_id) ON DELETE CASCADE ON UPDATE NO ACTION
 );

CREATE INDEX _reports_fkey_user_id ON _reports(_user_id);

CREATE TABLE _report_entries (
    _id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(), 
    _created_at TIMESTAMP NOT NULL DEFAULT NOW(), 
    _updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    _amount SMALLINT NOT NULL CHECK(_amount > 0),
    _carbohydrates NUMERIC(4,2) NOT NULL DEFAULT 0.00 CHECK(_carbohydrates > -1), 
    _proteins NUMERIC(4,2) NOT NULL DEFAULT 0.00 CHECK(_proteins > -1),
    _fats NUMERIC(4,2) NOT NULL DEFAULT 0.00 CHECK(_fats > -1),
    _product_id UUID NOT NULL,
    _report_id UUID NOT NULL,
    FOREIGN KEY(_product_id) REFERENCES _products(_id) ON DELETE CASCADE ON UPDATE NO ACTION,
    FOREIGN KEY(_report_id) REFERENCES _reports(_id) ON DELETE CASCADE ON UPDATE NO ACTION
);

CREATE INDEX _report_entries_fkey_report_id ON _report_entries(_report_id);
-- +goose Down

DROP TABLE _report_entries;

DROP TABLE _reports;

DROP TABLE _products;

DROP TABLE _daily_limit;

DROP TABLE _vault;

DROP TABLE _users;

DROP EXTENSION "uuid-ossp";