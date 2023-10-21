-- +goose Up
CREATE EXTENSION "uuid-ossp";

CREATE TABLE _users (
    _id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(), 
    _created_at TIMESTAMP NOT NULL DEFAULT NOW(), 
    _updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    _last_login TIMESTAMP NOT NULL DEFAULT NOW(),
    _active BOOLEAN NOT NULL DEFAULT TRUE,
    _name TEXT NOT NULL CHECK(CHAR_LENGTH(_name) > 0),
    _email TEXT NOT NULL UNIQUE CHECK(CHAR_LENGTH(_name) > 0),
    _phone_number TEXT NOT NULL CHECK(CHAR_LENGTH(_name) > 0)
);

CREATE TABLE _vault (
    _id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(), 
    _created_at TIMESTAMP NOT NULL DEFAULT NOW(), 
    _updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    _hashed_pw CHAR(64) NOT NULL,
    _user_id UUID NOT NULL UNIQUE,
    FOREIGN KEY (_user_id) REFERENCES _users(_id) ON DELETE CASCADE ON UPDATE NO ACTION
);

CREATE TABLE _products (
    _id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(), 
    _created_at TIMESTAMP NOT NULL DEFAULT NOW(), 
    _updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    _name TEXT NOT NULL UNIQUE CHECK(CHAR_LENGTH(_name) > 0),
    _amount SMALLINT NOT NULL CHECK(_amount > 0),
    _carbohydrate SMALLINT NOT NULL DEFAULT 0 CHECK(_carbohydrate > -1), 
    _protein SMALLINT NOT NULL DEFAULT 0 CHECK(_protein > -1),
    _fat SMALLINT NOT NULL DEFAULT 0 CHECK(_fat > -1),
    _user_id UUID NOT NULL,
    FOREIGN KEY (_user_id) REFERENCES _users(_id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE _reports (
    _id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    _created_at TIMESTAMP NOT NULL DEFAULT NOW(), 
    _updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    _date DATE NOT NULL UNIQUE,
    _amout_of_entries SMALLINT NOT NULL CHECK(_amout_of_entries > 0),
    _carbohydrates SMALLINT NOT NULL DEFAULT 0 CHECK(_carbohydrates > -1), 
    _proteins SMALLINT NOT NULL DEFAULT 0 CHECK(_proteins > -1),
    _fats SMALLINT NOT NULL DEFAULT 0 CHECK(_fats > -1),
    _user_id UUID NOT NULL,
    FOREIGN KEY (_user_id) REFERENCES _users(_id) ON DELETE CASCADE ON UPDATE CASCADE
 );

CREATE TABLE _report_entries (
    _id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(), 
    _created_at TIMESTAMP NOT NULL DEFAULT NOW(), 
    _updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    _amount SMALLINT NOT NULL CHECK(_amount > 0),
    _carbohydrates SMALLINT NOT NULL DEFAULT 0 CHECK(_carbohydrates > -1), 
    _proteins SMALLINT NOT NULL DEFAULT 0 CHECK(_proteins > -1),
    _fats SMALLINT NOT NULL DEFAULT 0 CHECK(_fats > -1),
    _product_id UUID NOT NULL,
    _report_id UUID NOT NULL,
    FOREIGN KEY(_product_id) REFERENCES _products(_id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY(_report_id) REFERENCES _reports(_id) ON DELETE NO ACTION ON UPDATE NO ACTION
);

-- +goose Down

DROP TABLE _report_entries;

DROP TABLE _reports;

DROP TABLE _products;

DROP TABLE _vault;

DROP TABLE _users;

DROP EXTENSION "uuid-ossp";