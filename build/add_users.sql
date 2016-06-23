USE `httpdb`;

INSERT INTO `userpasswd`(user_id,name,password ) VALUES (NULL, 'admin', 'admin'),(NULL, 'jiawei.zhu', '123456');

INSERT INTO `userns`(user_id, namespace) (SELECT `user_id`, 'admin_namespace1' from `userpasswd` WHERE `name`='admin');

INSERT INTO `userns`(user_id, namespace) (SELECT `user_id`, 'admin_namespace2' from `userpasswd` WHERE `name`='admin');

INSERT INTO `userns`(user_id, namespace) (SELECT `user_id`, 'jiawei.zhu_namespace1' from `userpasswd` WHERE `name`='jiawei.zhu');

INSERT INTO `userns`(user_id, namespace) (SELECT `user_id`, 'jiawei.zhu_namespace2' from `userpasswd` WHERE `name`='jiawei.zhu');