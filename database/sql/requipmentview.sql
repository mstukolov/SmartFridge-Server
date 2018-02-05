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