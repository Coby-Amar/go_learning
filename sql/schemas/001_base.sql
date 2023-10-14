-- +goose Up
CREATE EXTENSION "uuid-ossp";
CREATE EXTENSION "moddatetime";

CREATE TABLE _products (
    _id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(), 
    _created_at TIMESTAMP NOT NULL DEFAULT NOW(), 
    _updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    _name VARCHAR(200) NOT NULL UNIQUE CHECK(CHAR_LENGTH(_name) > 0),
    _amount SMALLINT NOT NULL CHECK(_amount > 0),
    _carbohydrate SMALLINT NOT NULL DEFAULT 0, 
    _protein SMALLINT NOT NULL DEFAULT 0,
    _fat SMALLINT NOT NULL DEFAULT 0
);
CREATE TRIGGER mdt_moddatetime
    BEFORE UPDATE ON _products
    FOR EACH ROW
    EXECUTE PROCEDURE moddatetime(moddate);

CREATE TABLE _reports (
    _id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(), 
    _created_at TIMESTAMP NOT NULL DEFAULT NOW(), 
    _updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    _date DATE NOT NULL UNIQUE
);
CREATE TRIGGER mdt_moddatetime
    BEFORE UPDATE ON _reports
    FOR EACH ROW
    EXECUTE PROCEDURE moddatetime(moddate);

CREATE TABLE _report_entries (
    _id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(), 
    _created_at TIMESTAMP NOT NULL DEFAULT NOW(), 
    _updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    _product_id UUID NOT NULL REFERENCES _products(_id) ON DELETE CASCADE,
    _report_id UUID NOT NULL REFERENCES _reports(_id) ON DELETE CASCADE,
    _amount SMALLINT NOT NULL CHECK(_amount > 0),
    _carbohydrate SMALLINT NOT NULL DEFAULT 0, 
    _protein SMALLINT NOT NULL DEFAULT 0,
    _fat SMALLINT NOT NULL DEFAULT 0
);

CREATE TRIGGER mdt_moddatetime
    BEFORE UPDATE ON _report_entries
    FOR EACH ROW
    EXECUTE PROCEDURE moddatetime(moddate);

-- +goose Down
DROP TRIGGER mdt_moddatetime ON _report_entries;
DROP TABLE _report_entries;

DROP TRIGGER mdt_moddatetime ON _reports;
DROP TABLE _reports;

DROP TRIGGER mdt_moddatetime ON _products;
DROP TABLE _products;

DROP EXTENSION "uuid-ossp";
DROP EXTENSION "moddatetime";