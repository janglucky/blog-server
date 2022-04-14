create database  if not exists tensir_blog;


create table if not exists tensir_blog_user (
                                                id int auto_increment,
                                                username varchar(60),
    password varchar(60),
    role varchar(50),

    )