-- DROPS
DROP TABLE drops;

CREATE TABLE drops (
  Id              serial primary key,
  Name            varchar(255),
  Latitude        float(8) not null,
  Longitude       float(8) not null,
  Radius          integer not null,
  DropGeom        geometry,
  CreatedAt       timestamp not null default now(),
  UpdatedAt       timestamp not null default now()
);

CREATE OR REPLACE FUNCTION build_geom() RETURNS trigger AS $build_geom_trigger$
  BEGIN
    -- NEW.DropGeom := ST_GeomFromText('POINT(' || NEW.Longitude || ' ' || NEW.Latitude || ')', 4326);
    NEW.DropGeom := ST_Buffer(ST_SetSRID(ST_MakePoint(NEW.Longitude, NEW.Latitude), 4326)::geography, NEW.Radius)::geometry;
    RETURN NEW;
  END;
$build_geom_trigger$ LANGUAGE plpgsql;

CREATE TRIGGER build_geom_trigger BEFORE INSERT ON drops FOR EACH ROW EXECUTE PROCEDURE build_geom();

