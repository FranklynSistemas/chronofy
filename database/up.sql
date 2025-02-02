DROP TABLE IF EXISTS events;

CREATE TABLE events (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  external_id VARCHAR(32) NOT NULL,
  payload JSONB NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

INSERT INTO events (name, external_id, payload) VALUES ('event_type1', 'external_id1', '{"key1": "value1"}');
INSERT INTO events (name, external_id, payload) VALUES ('event_type2', 'external_id2', '{"key2": "value2"}');
INSERT INTO events (name, external_id, payload) VALUES ('event_type3', 'external_id3', '{"key3": "value3"}');
INSERT INTO events (name, external_id, payload) VALUES ('event_type4', 'external_id4', '{"key4": "value4"}');