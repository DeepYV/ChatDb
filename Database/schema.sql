

CREATE TABLE IF NOT EXISTS chatroom
(
    chatroomname text COLLATE pg_catalog."default" NOT NULL,
    createdat timestamp without time zone,
    member integer DEFAULT 0,
    CONSTRAINT chatroom_pkey PRIMARY KEY (chatroomname)
)

CREATE TABLE IF NOT EXISTS joined
(
    user_id text ,
    createdat timestamp without time zone,
    chatroomname text ,
    role text ,
    join_id text,
    CONSTRAINT joined_pkey PRIMARY KEY (join_id),
    CONSTRAINT joined_chatroomname_fkey FOREIGN KEY (chatroomname)
        REFERENCES chatroom (chatroomname) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
        NOT VALID
)

CREATE TABLE IF NOT EXISTS message
(
    chatroomname text ,
    message text ,
    user_id text ,
    createdat timestamp without time zone,
    CONSTRAINT message_chatroomname_fkey FOREIGN KEY (chatroomname)
        REFERENCES chatroom (chatroomname) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
        NOT VALID,
    CONSTRAINT message_user_id_fkey FOREIGN KEY (user_id)
        REFERENCES user (user_id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
        NOT VALID
)

CREATE TABLE IF NOT EXISTS user
(
    user_id text,
    username text ,
    password text ,
    online text ,
    createat timestamp without time zone,
    CONSTRAINT "User_pkey" PRIMARY KEY (user_id)
)