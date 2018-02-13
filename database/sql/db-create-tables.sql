--create customers table---
CREATE TABLE customers
(
  id        BIGSERIAL NOT NULL
    CONSTRAINT customers_id_pk
    PRIMARY KEY,
  name      VARCHAR(150),
  createdat TIMESTAMP DEFAULT now(),
  updatedat TIMESTAMP DEFAULT now()
);

COMMENT ON TABLE customers IS 'Клиенты';

--create emploees table---
CREATE TABLE emploees
(
  id        SERIAL NOT NULL
    CONSTRAINT users_pkey
    PRIMARY KEY,
  login     VARCHAR(100),
  password  VARCHAR(100),
  createdat TIMESTAMP DEFAULT now(),
  updatedat TIMESTAMP DEFAULT now(),
  surname   VARCHAR(150),
  name      VARCHAR(150),
  hash      VARCHAR(100),
  auth      BOOLEAN   DEFAULT TRUE
);

CREATE UNIQUE INDEX users_id_uindex
  ON emploees (id);

--create microcontrollers table---
CREATE TABLE locationequipment
(
  id        BIGSERIAL NOT NULL
    CONSTRAINT locationequipment_id_pk
    PRIMARY KEY,
  address   VARCHAR(500),
  lat       NUMERIC,
  lng       NUMERIC,
  updatedat TIMESTAMP DEFAULT now(),
  createdat TIMESTAMP DEFAULT now()
);
COMMENT ON TABLE locationequipment IS 'расположение оборудования';


--create microcontrollers table---
CREATE TABLE microcontrollers
(
  id           BIGSERIAL    NOT NULL
    CONSTRAINT " microcontrollers_pkey"
    PRIMARY KEY,
  deviceid     VARCHAR(100) NOT NULL,
  requipmentid BIGINT       NOT NULL,
  emptyweight  DOUBLE PRECISION,
  fullweight   DOUBLE PRECISION,
  factor       DOUBLE PRECISION
);

CREATE UNIQUE INDEX " microcontrollers_id_uindex"
  ON microcontrollers (id);

CREATE UNIQUE INDEX " microcontrollers_deviceid_uindex"
  ON microcontrollers (deviceid);

CREATE UNIQUE INDEX " microcontrollers_requipmentid_uindex"
  ON microcontrollers (requipmentid);
--create requipmentlasttrans table---
CREATE TABLE requipmentlasttrans
(
  id                BIGSERIAL NOT NULL
    CONSTRAINT requipmentlasttrans_id_pk
    PRIMARY KEY,
  retailequipmentid BIGINT,
  sentsortypeid     BIGINT,
  sensorvalue       INTEGER,
  createdat         TIMESTAMP DEFAULT now(),
  updatedat         TIMESTAMP DEFAULT now(),
  fullness          DOUBLE PRECISION
);

--create requipmenttrans table---
CREATE TABLE requipmenttrans
(
  id                BIGSERIAL NOT NULL
    CONSTRAINT requipmenttrans_pkey
    PRIMARY KEY,
  retailequipmentid BIGINT,
  sentsortypeid     BIGINT,
  sensorvalue       INTEGER,
  createdat         TIMESTAMP DEFAULT now(),
  updatedat         TIMESTAMP DEFAULT now(),
  fullness          DOUBLE PRECISION
);

CREATE UNIQUE INDEX requipmenttrans_id_uindex
  ON requipmenttrans (id);
COMMENT ON TABLE requipmenttrans IS 'все сообщения';
COMMENT ON TABLE requipmentlasttrans IS 'все сообщения';

--create retailchains table---
CREATE TABLE retailchains
(
  id         BIGSERIAL NOT NULL
    CONSTRAINT retailchain_pkey
    PRIMARY KEY,
  name       VARCHAR(200),
  createdat  TIMESTAMP DEFAULT now(),
  updatedat  TIMESTAMP DEFAULT now(),
  customerid INTEGER
);

CREATE INDEX retailchain_id_name_index
  ON retailchains (id, name);


--create retailequipments table---
CREATE TABLE retailequipments
(
  id                  BIGSERIAL NOT NULL
    CONSTRAINT retailequipment_id_pk
    PRIMARY KEY,
  createdat           TIMESTAMP DEFAULT now(),
  updatedat           TIMESTAMP DEFAULT now(),
  maxvalue            DOUBLE PRECISION,
  serialnumber        VARCHAR(50),
  locationequipmentid BIGINT,
  lastvalue           DOUBLE PRECISION,
  filling             DOUBLE PRECISION,
  retailstore_id      BIGINT
);
COMMENT ON TABLE retailequipments IS 'Торговое оборудование';


--create retailequipments table---
CREATE TABLE retailstores
(
  id             BIGSERIAL NOT NULL,
  name           VARCHAR(150),
  createdat      TIMESTAMP DEFAULT now(),
  updatedat      TIMESTAMP DEFAULT now(),
  retailchain_id BIGINT
);
COMMENT ON TABLE retailstores IS 'Торговая точка';

--create requipmentdetailviews view---
CREATE VIEW requipmentdetailviews AS
  SELECT eqiup.id AS requipid,
    eqiup.serialnumber AS requipserialnumber,
    eqiup.maxvalue AS requipmaxvalue,
    trans.fullness AS requipfullness,
    trans.sensorvalue,
    trans.createdat,
    store.id AS storeid,
    store.name AS storename,
    chain.id AS rchainid,
    chain.name AS rchainname,
    location.address,
    location.lng,
    location.lat
   FROM ((((retailequipments eqiup
     JOIN retailstores store ON ((eqiup.retailstore_id = store.id)))
     JOIN locationequipment location ON ((eqiup.locationequipmentid = location.id)))
     JOIN retailchains chain ON ((store.retailchain_id = chain.id)))
     LEFT JOIN requipmentlasttrans trans ON ((eqiup.id = trans.retailequipmentid)));

--create retailequipmentgps view---
CREATE VIEW retailequipmentgps AS
  SELECT eq.id,
    eq.serialnumber,
    loc.address,
    loc.lat,
    loc.lng
   FROM (((retailequipments eq
     JOIN retailstores st ON ((eq.retailstore_id = st.id)))
     JOIN retailchains ch ON ((st.retailchainid = ch.id)))
     JOIN locationequipment loc ON ((eq.locationequipmentid = loc.id)));

CREATE VIEW requipmentview AS
select eqiup.id as requipid,
        eqiup.serialnumber as requipserialnumber,
          eqiup.filling as requipfilling,
            eqiup.maxvalue as requipmaxvalue,
            eqiup.lastvalue as requiplastvalue,
            store.id as storeid,
              store.name as storename,
                chain.id as rchainid,
                  chain.name as rchainname

from retailequipments as eqiup
  join retailstores store on eqiup.retailstore_id = store.id
    join retailchains chain on store.retailchain_id = chain.id
;

--create requipmentviews view---
CREATE VIEW requipmentviews AS
SELECT eqiup.id AS requipid,
       eqiup.serialnumber AS requipserialnumber,
       -100*trans.fullness AS requipfullness,
        (trans.sensorvalue - micro.emptyweight)/micro.factor as sensorvalue,
        trans.createdat as createdat,
        store.id AS storeid,
        store.name AS storename,
        chain.id AS rchainid,
        chain.name AS rchainname
FROM ((((retailequipments eqiup
  JOIN retailstores store ON ((eqiup.retailstore_id = store.id)))
  JOIN locationequipment location ON ((eqiup.locationequipmentid = location.id)))
  JOIN retailchains chain ON ((store.retailchain_id = chain.id)))
  join microcontrollers micro ON  eqiup.id = micro.requipmentid
  LEFT JOIN requipmentlasttrans trans ON ((eqiup.id = trans.retailequipmentid)));