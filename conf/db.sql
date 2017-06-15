CREATE OR REPLACE VIEW v_friends AS
 SELECT u1.id AS user1id,
    u1.email AS user1email,
    u2.id AS user2id,
    u2.email AS user2email
   FROM api_friends f
     LEFT JOIN api_user u1 ON f.userid1 = u1.id
     LEFT JOIN api_user u2 ON f.userid2 = u2.id;
	
CREATE OR REPLACE VIEW v_permission AS
 SELECT u1.id AS requestor,
    u1.email AS requestoremail,
    u2.id AS target,
    u2.email AS targetemail,
    p.status 
   FROM api_permission p
     LEFT JOIN api_user u1 ON p.requestor = u1.id
     LEFT JOIN api_user u2 ON p.target = u2.id;	