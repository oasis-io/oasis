
USE oasis;

#用户表
insert into users(username,password) values("admin","123456");
insert into users(username,password) values("zhangshaodong","123456");

#角色表
INSERT INTO user_roles(name) VALUES ("Admin"), ("DBA");
INSERT INTO user_role_relation(user_role_id,user_id) VALUES (1,1), (2,2);


#菜单表
INSERT INTO menus (parent_id, sort, name, path, component, title, icon ) VALUES (0, 0, "Home", "home", "views/home/index.vue", "首页", "HomeFilled");

INSERT INTO menus (parent_id, sort, name, path, component, title, icon ) VALUES (0, 0, "Instance", "instance","views/instance/index.vue", "实例列表",  "Menu");

INSERT INTO menus (parent_id, sort, name, path, component, title, icon ) VALUES (0, 0, "User", "user", "views/user/index.vue", "用户中心", "User");

INSERT INTO menus (parent_id, sort, name, path, component, title ) VALUES (3, 1, "UserList", "list", "views/user/UserList/index.vue", "用户管理");

INSERT INTO menus (parent_id, sort, name, path, component, title ) VALUES (3, 2, "UserRole", "role", "views/user/UserRole/index.vue", "角色管理");

INSERT INTO menus (parent_id, sort, name, path, component, title ) VALUES (3, 3, "UserGroup", "group", "views/user/UserGroup/index.vue", "用户组管理");


#角色菜单表关联数据
INSERT INTO role_menu_relations(role_id,menu_id) VALUES(1,1), (2,1);