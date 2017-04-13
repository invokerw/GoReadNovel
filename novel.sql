-- MySQL dump 10.14  Distrib 5.5.52-MariaDB, for Linux (x86_64)
--
-- Host: localhost    Database: novel
-- ------------------------------------------------------
-- Server version	5.5.52-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `novel`
--

DROP TABLE IF EXISTS `novel`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `novel` (
  `novelid` int(11) NOT NULL AUTO_INCREMENT COMMENT '//小说ID',
  `name` varchar(50) NOT NULL COMMENT '//小说名称',
  `author` varchar(30) DEFAULT '未知' COMMENT '//作者',
  `noveldesc` varchar(300) DEFAULT '暂无' COMMENT '//小说描述',
  `noveltype` varchar(30) DEFAULT '其他' COMMENT '//小说类型',
  `addr` varchar(80) NOT NULL COMMENT '//小说地址',
  `imageaddr` varchar(80) DEFAULT NULL COMMENT '//图片地址',
  `lchaptername` varchar(70) DEFAULT NULL COMMENT '//最新章节名称',
  `lchapteraddr` varchar(80) DEFAULT NULL COMMENT '//最新章节地址',
  `status` varchar(30) DEFAULT '连载中' COMMENT '//连载还是完结',
  PRIMARY KEY (`novelid`),
  UNIQUE KEY `name` (`name`) USING BTREE,
  KEY `author` (`author`),
  KEY `noveltype` (`noveltype`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `novel`
--

LOCK TABLES `novel` WRITE;
/*!40000 ALTER TABLE `novel` DISABLE KEYS */;
INSERT INTO `novel` VALUES (10,'圣墟','辰东','    在破败中崛起，在寂灭中复苏。\r\n    沧海成尘，雷电枯竭，那一缕幽雾又一次临近大地，世间的枷锁被打开了，一个全新的世界就此揭开神秘的一角……\r\n各位书友要是觉得《圣墟》还不…','其他','http://www.huanyue123.com/book/0/11/','http://www.huanyue123.com/files/article/image/0/11/11s.jpg','第三百三十三章 太上八卦炉','/book/0/11/7740143.html','连载中'),(11,'一念永恒','耳根','    一念成沧海，一念化桑田。一念斩千魔，一念诛万仙。唯我念……永恒\r\n各位书友要是觉得《一念永恒》还不错的话请不要忘记向您QQ群和微博里的朋友推荐哦！\r\n\r\n…','其他','http://www.huanyue123.com/book/0/4/','http://www.huanyue123.com/files/article/image/0/4/4s.jpg','第718章 大家别冲动','/book/0/4/7727488.html','连载中'),(12,'校花的贴身高手','鱼人二代','    一个大山里走出来的绝世高手，一块能预知未来的神秘玉佩……林逸是一名普通学生，不过，他还身负另外一个重任，那就是追校花！而且还是奉校花老爸之命！虽然林逸很不想跟这位难伺…','其他','http://www.huanyue123.com/book/0/2/','http://www.huanyue123.com/files/article/image/0/2/2s.jpg','第6161章 得到传承','/book/0/2/7747344.html','连载中'),(13,'大主宰','天蚕土豆','    大千世界，位面交汇，万族林立，群雄荟萃，一位位来自下位面的天之至尊，在这无尽世界，演绎着令人向往的传奇，追求着那主宰之路。\r\n    无尽火域，炎帝执掌，万火焚苍穹。\r\n    …','其他','http://www.huanyue123.com/book/0/5/','http://www.huanyue123.com/files/article/image/0/5/5s.jpg','第一千四百七十二章 秦东海','/book/0/5/7726351.html','连载中'),(14,'莽荒纪','我吃西红柿','一部小说，就是一个世界。 在《莽荒纪》这个世界里，修炼者为了生存，和天斗，和地斗，和妖斗的部落人们。 有为了逍遥长生，历三灾九劫，纵死无悔的修仙者。 更有夸父逐日…… 更有后…','其他','http://www.huanyue123.com/book/0/1/','http://www.huanyue123.com/files/article/image/0/1/1s.jpg','番茄新书《雪鹰领主》正式上传！','/book/0/1/329882.html','连载中'),(15,'玄界之门','忘语','    天降神物！异血附体！\r\n    群仙惊惧！万魔退避！\r\n    一名从东洲大陆走出的少年。\r\n    一具生死相依的红粉骷髅。\r\n    一个立志成为至强者的故事。\r\n    一段叱咤星河，大闹三…','其他','http://www.huanyue123.com/book/0/6/','http://www.huanyue123.com/files/article/image/0/6/6s.jpg','第九百三十六章 血斗','/book/0/6/7732952.html','连载中'),(16,'逆鳞','柳下挥','    天生废材，遭遇龙神附体。继承了神龙的意念和能力，生鳞幻爪、御水龙息、行云降雨，肉身无敌。在这个人人都想屠龙的时代，李牧羊一直生活的很有压力。\r\n各位书友要是觉得《逆鳞》…','其他','http://www.huanyue123.com/book/0/299/','http://www.huanyue123.com/files/article/image/0/299/299s.jpg','第六百七十一章、白云长老！','/book/0/299/7610905.html','连载中'),(17,'巫神纪','血红','    当历史变成传说\r\n    当传说变成神话\r\n    当神话都已经斑驳点点\r\n    当时间的沙尘湮没一切\r\n    我们的名字，我们的故事，依旧在岁月的长河中传播\r\n    一如太阳高悬天空，永恒…','其他','http://www.huanyue123.com/book/0/3/','http://www.huanyue123.com/files/article/image/0/3/3s.jpg','感言及预告！','/book/0/3/5123904.html','连载中'),(18,'英雄联盟：上帝之手','三千勿忘尽','    作为国内最强战队的队长，电竞豪门的绝对王牌，他将自己最好的青春全部奉献给了战队，然而，当英雄迟暮，等待他的，却是被无情的扫地出门！\r\n    这一世，当忠诚化为可笑的悲哀，…','其他','http://www.huanyue123.com/book/0/236/','http://www.huanyue123.com/files/article/image/0/236/236s.jpg','第三百六十四章 终结！（第二更）','/book/0/236/7611851.html','连载中'),(19,'太古神王','净无痕','    九天大陆，天穹之上有九条星河，亿万星辰，皆为武命星辰，武道之人，可沟通星辰，觉醒星魂，成武命修士。\r\n    传说，九天大陆最为厉害的武修，每突破一个境界，便能开辟一扇星门…','其他','http://www.huanyue123.com/book/1/1659/','http://www.huanyue123.com/files/article/image/1/1659/1659s.jpg','第1642章 不是秦问天？（五更）','/book/1/1659/7752584.html','连载中'),(20,'雪鹰领主','我吃西红柿','    深海魔兽的呼吸形成永不停息的风暴…\r\n    熔岩巨人的脚步毁灭一座座城池…\r\n    深渊恶魔想要侵入这座世界…\r\n    而神灵降临，行走人间传播他的光辉…\r\n    然而整个世界由龙山…','其他','http://www.huanyue123.com/book/0/8/','http://www.huanyue123.com/files/article/image/0/8/8s.jpg','第36篇 第19章 虚界幻境道第一杀招（下）','/book/0/8/7612825.html','连载中'),(21,'武炼巅峰','莫默','    武之巅峰，是孤独，是寂寞，是漫漫求索，是高处不胜寒\r\n    逆境中成长，绝地里求生，不屈不饶，才能堪破武之极道。\r\n    凌霄阁试炼弟子兼扫地小厮杨开偶获一本无字黑书，从此踏…','其他','http://www.huanyue123.com/book/0/45/','http://www.huanyue123.com/files/article/image/0/45/45s.jpg','第三千五百四十二章 出尔反尔','/book/0/45/7728745.html','连载中'),(22,'天域苍穹','风凌天下','    笑尽天下英雄，宇内我为君主！万水千山，以我为尊；八荒六合，唯我称雄！我欲舞风云，凌天下，踏天域，登苍穹！谁可争锋？！诸君可愿陪我，并肩凌天下，琼霄风云舞，征战这天域苍…','其他','http://www.huanyue123.com/book/0/28/','http://www.huanyue123.com/files/article/image/0/28/28s.jpg','第2037章 大结局！','/book/0/28/5534024.html','连载中'),(23,'永夜君王','烟雨江南','    千夜自困苦中崛起，在背叛中坠落。自此一个人，一把枪，行在永夜与黎明之间，却走出一段传奇。若永夜注定是他的命运，那他也要成为主宰的王。\r\n各位书友要是觉得《永夜君王》还不…','其他','http://www.huanyue123.com/book/0/49/','http://www.huanyue123.com/files/article/image/0/49/49s.jpg','章二四六 天赋和智慧','/book/0/49/7502460.html','连载中'),(24,'飞天','跃千愁','    来历神秘，风华绝代的美男子，回眸间睥睨天下，却在诡谲莫测的‘万丈红尘’大阵之中，守一只惊心动魄的断弦古琴，静候有缘人！\r\n    小子得遇，方知，苍穹之下世态炎凉，妖魔鬼怪…','其他','http://www.huanyue123.com/book/0/42/','http://www.huanyue123.com/files/article/image/0/42/42s.jpg','完本感言','/book/0/42/4974390.html','连载中'),(25,'龙王传说','唐家三少','    伴随着魂导科技的进步，斗罗大6上的人类征服了海洋，又现了两片大6。魂兽也随着人类魂师的猎杀无度走向灭亡，沉睡无数年的魂兽之王在星斗大森林最后的净土苏醒，它要带领仅存的族…','其他','http://www.huanyue123.com/book/0/7/','http://www.huanyue123.com/files/article/image/0/7/7s.jpg','第九百七十九章 练枪','/book/0/7/7740280.html','连载中'),(26,'不朽凡人','鹅是老五','    我，只有凡根，一介凡人！\r\n    我，叫莫无忌！\r\n    我，要不朽！\r\n…','其他','http://www.huanyue123.com/book/0/10/','http://www.huanyue123.com/files/article/image/0/10/10s.jpg','第七百一十五章 强大的仙尊初期','/book/0/10/7747478.html','连载中'),(27,'黑卡','萧瑟良','    一张神秘的黑卡，每周都会放不同的额度，石磊必须在一周时间内将所有额度消费完毕，否则，将迎接黑卡的惩罚。\n    “花钱真的是个体力活”——石磊如是说。\n    “先定一个小目标…','其他','http://www.huanyue123.com/book/2/2749/','http://www.huanyue123.com/files/article/image/2/2749/2749s.jpg','第五百八十一章 抽奖结果','/book/2/2749/7717660.html','连载中'),(28,'邪王，宠不够！','灵妖妖','    圣医后人穿成将军府弃女，却被未来‘姐夫’给睡了！\r\n    敢睡她？她便睡回去！\r\n    只是这位王爷，你不是病入膏肓么？那为啥夜夜龙精虎猛，总让她脚软跑不动？\r\n    听过，某男…','其他','http://www.huanyue123.com/book/5/5642/','http://www.huanyue123.com/files/article/image/5/5642/5642s.jpg','第553章 穷追不舍','/book/5/5642/7611384.html','连载中'),(29,'道君','跃千愁','   一个地球神级盗墓宗师，闯入修真界的故事……\r\n　　桃花源里，有歌声。\r\n　　山外青山，白骨山。\r\n　　五花马，千金裘，倚天剑，应我多情，啾啾鬼鸣，美人薄嗔。\r\n　　天地无垠，谁…','其他','http://www.huanyue123.com/book/5/5544/','http://www.huanyue123.com/files/article/image/5/5544/5544s.jpg','第一二四章 想干什么?','/book/5/5544/7614072.html','连载中'),(30,'惊悚乐园','三天两觉','    欢迎来到惊悚乐园。\n    这不仅是游戏，也是挑战和试炼。\n    恐惧是人类的本能，它使人软弱，惊慌，从而犯下错误。\n    金钱可以将游戏者武装起来，但智慧和勇气是买不到的。\n  …','其他','http://www.huanyue123.com/book/0/939/','http://www.huanyue123.com/files/article/image/0/939/939s.jpg','第十六章 敌军已抵达战场','/book/0/939/7759400.html','连载中'),(31,'武动乾坤','天蚕土豆','    修炼一途，乃窃阴阳，夺造化，转涅盘，握生死，掌轮回。\r\n    武之极，破苍穹，动乾坤！\r\n    各位书友要是觉得《武动乾坤》还不错的话请不要忘记向您QQ群和微博里的朋友推荐哦！…','其他','http://www.huanyue123.com/book/0/36/','http://www.huanyue123.com/files/article/image/0/36/36s.jpg','大结局活动，1744，欢迎大家。','/book/0/36/6626782.html','连载中'),(32,'斗破苍穹','天蚕土豆','    这里是属于斗气的世界，没有花俏艳丽的魔法，有的，仅仅是繁衍到巅峰的斗气！\n    想要知道异界的斗气在发展到巅峰之后是何种境地吗？请观斗破苍穹^_^\n    PS:据调查，斗气，并非…','其他','http://www.huanyue123.com/book/0/35/','http://www.huanyue123.com/files/article/image/0/35/35s.jpg','第一章 五帝破空','/book/0/35/335873.html','连载中'),(33,'帝霸','厌笔萧生','    千万年前，李七夜栽下一株翠竹。\n    八百万年前，李七夜养了一条鲤鱼。\n    五百万年前，李七夜收养一个小女孩。\n    …………………………\n    今天，李七夜一觉醒来，翠竹修练…','其他','http://www.huanyue123.com/book/0/54/','http://www.huanyue123.com/files/article/image/0/54/54s.jpg','第2425章巨鹰','/book/0/54/7729917.html','连载中'),(34,'斗罗大陆','唐家三少','    唐门外门弟子唐三，因偷学内门绝学为唐门所不容，跳崖明志时却来到了另一个世界，一个属于武魂的世界。名叫斗罗大陆。\n    这里没有魔法，没有斗气，没有武术，却有神奇的武魂。这…','其他','http://www.huanyue123.com/book/0/210/','http://www.huanyue123.com/files/article/image/0/210/210s.jpg','第二百三十六章 大结局，最后一个条件（全书完）','/book/0/210/151345.html','连载中'),(35,'通天仙路','苍天白鹤','    “欧阳大师，求您给我锻造个神器吧！条件您开！”\n    “……”\n    “欧阳大师，我知道全天下就您能锻造七属性的神器，请您成全！”\n    “……”\n    “欧阳大师，您……”\n   …','其他','http://www.huanyue123.com/book/2/2609/','http://www.huanyue123.com/files/article/image/2/2609/2609s.jpg','第三百九十一章 新的心得','/book/2/2609/7737206.html','连载中'),(36,'儒道至圣','永恒之火','    这是一个读书人掌握天地之力的世界。\r\n    才气在身，诗可杀敌，词能灭军，文章安天下。\r\n    秀才提笔，纸上谈兵；举人杀敌，出口成章；进士一怒，唇枪舌剑。\r\n    圣人驾临，口…','其他','http://www.huanyue123.com/book/0/412/','http://www.huanyue123.com/files/article/image/0/412/412s.jpg','第2082章 山脉圣灵','/book/0/412/7607108.html','连载中'),(37,'超品相师','九灯和善','    相师分九品，一品一重天\n    风水有境界，明理，养气，修身，问道。\n    二十一世纪的一位普通青年偶获诸葛亮生前的玄学传承，从此混迹都市，游走于高官权贵之间，豪门千金，世家…','其他','http://www.huanyue123.com/book/4/4355/','http://www.huanyue123.com/files/article/image/4/4355/4355s.jpg','第2845章 以一敌六','/book/4/4355/7742912.html','连载中'),(38,'择天记','猫腻','    太始元年，有神石自太空飞来，分散落在人间，其中落在东土大陆的神石，上面镌刻着奇怪的图腾，人因观其图腾而悟道，后立国教。\n    数千年后，十四岁的少年孤儿陈长生，为治病改命…','其他','http://www.huanyue123.com/book/0/30/','http://www.huanyue123.com/files/article/image/0/30/30s.jpg','第1164章 最后的晚餐以及谈话','/book/0/30/7741519.html','连载中'),(39,'容你轻轻撩动我心','蓝岚','    苏暮然从未想到，和上司捉未婚妻的奸，奸夫居然是她男朋友。“既然他们玩的很开心，不如，我们也凑合吧！”门外，上司一张俊脸冷若冰霜，却突然扭过头对她一本正经道。苏暮然被惊…','其他','http://www.huanyue123.com/book/5/5346/','http://www.huanyue123.com/files/article/image/5/5346/5346s.jpg','第243章 吾家有儿初长成','/book/5/5346/7715012.html','连载中');
/*!40000 ALTER TABLE `novel` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2017-04-13 23:33:29
