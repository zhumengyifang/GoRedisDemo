package internal

import "github.com/garyburd/redigo/redis"

//http://www.redis.cn/commands/zincrby.html

/**
将所有指定成员添加到键为key有序集合（sorted set）里面。 添加时可以指定多个分数/成员（score/member）对。 如果指定添加的成员已经是有序集合里面的成员，则会更新改成员的分数（scrore）并更新到正确的排序位置。

如果key不存在，将会创建一个新的有序集合（sorted set）并将分数/成员（score/member）对添加到有序集合，就像原来存在一个空的有序集合一样。如果key存在，但是类型不是有序集合，将会返回一个错误应答。

分数值是一个双精度的浮点型数字字符串。+inf和-inf都是有效值。
 */
func Zadd(conn redis.Conn) {

}

/**
返回key的有序集元素个数。
 */
func Zcard(conn redis.Conn) {

}

/**
返回有序集key中，score值在min和max之间(默认包括score值等于min或max)的成员。 关于参数min和max的详细使用方法，请参考ZRANGEBYSCORE命令。
 */
func Zcount(conn redis.Conn) {

}

/**
为有序集key的成员member的score值加上增量increment。如果key中不存在member，就在key中添加一个member，score是increment（就好像它之前的score是0.0）。如果key不存在，就创建一个只含有指定member成员的有序集合。

当key不是有序集类型时，返回一个错误。

score值必须是字符串表示的整数值或双精度浮点数，并且能接受double精度的浮点数。也有可能给一个负数来减少score的值。
 */
func Zincrby(conn redis.Conn) {

}

/**
计算给定的numkeys个有序集合的交集，并且把结果放到destination中。 在给定要计算的key和其它参数之前，必须先给定key个数(numberkeys)。

默认情况下，结果中一个元素的分数是有序集合中该元素分数之和，前提是该元素在这些有序集合中都存在。因为交集要求其成员必须是给定的每个有序集合中的成员，结果集中的每个元素的分数和输入的有序集合个数相等。

对于WEIGHTS和AGGREGATE参数的描述，参见命令ZUNIONSTORE。

如果destination存在，就把它覆盖。
 */
func Zinterstore(conn redis.Conn) {

}

/**
ZLEXCOUNT 命令用于计算有序集合中指定成员之间的成员数量。
 */
func Zlexcount(conn redis.Conn) {

}

/**
删除并返回有序集合key中的最多count个具有最高得分的成员。

如未指定，count的默认值为1。指定一个大于有序集合的基数的count不会产生错误。 当返回多个元素时候，得分最高的元素将是第一个元素，然后是分数较低的元素。
 */
func Zpopmax(conn redis.Conn) {

}

/**
删除并返回有序集合key中的最多count个具有最低得分的成员。

如未指定，count的默认值为1。指定一个大于有序集合的基数的count不会产生错误。 当返回多个元素时候，得分最低的元素将是第一个元素，然后是分数较高的元素。
 */
func Zpopmin(conn redis.Conn) {

}

/**
返回存储在有序集合key中的指定范围的元素。 返回的元素可以认为是按得分从最低到最高排列。 如果得分相同，将按字典排序。

当你需要元素从最高分到最低分排列时，请参阅ZREVRANGE（相同的得分将使用字典倒序排序）。

参数start和stop都是基于零的索引，即0是第一个元素，1是第二个元素，以此类推。 它们也可以是负数，表示从有序集合的末尾的偏移量，其中-1是有序集合的最后一个元素，-2是倒数第二个元素，等等。

start和stop都是全包含的区间，因此例如ZRANGE myzset 0 1将会返回有序集合的第一个和第二个元素。

超出范围的索引不会产生错误。 如果start参数的值大于有序集合中的最大索引，或者start > stop，将会返回一个空列表。 如果stop的值大于有序集合的末尾，Redis会将其视为有序集合的最后一个元素。

可以传递WITHSCORES选项，以便将元素的分数与元素一起返回。这样，返回的列表将包含value1,score1,...,valueN,scoreN，而不是value1,...,valueN。 客户端类库可以自由地返回更合适的数据类型（建议：具有值和得分的数组或记录）。
 */
func Zrange(conn redis.Conn) {

}

/**
ZRANGEBYLEX 返回指定成员区间内的成员，按成员字典正序排序, 分数必须相同。 在某些业务场景中,需要对一个字符串数组按名称的字典顺序进行排序时,可以使用Redis中SortSet这种数据结构来处理。
 */
func Zrangebylex(conn redis.Conn) {

}

/**

 */
func Zrangebyscore(conn redis.Conn) {

}

/**
返回有序集key中成员member的排名。其中有序集成员按score值递增(从小到大)顺序排列。排名以0为底，也就是说，score值最小的成员排名为0。

使用ZREVRANK命令可以获得成员按score值递减(从大到小)排列的排名。
 */
func Zrank(conn redis.Conn) {

}

/**
返回的是从有序集合中删除的成员个数，不包括不存在的成员。
 */
func Zrem(conn redis.Conn) {

}

/**
ZREMRANGEBYLEX 删除名称按字典由低到高排序成员之间所有成员。
不要在成员分数不同的有序集合中使用此命令, 因为它是基于分数一致的有序集合设计的,如果使用,会导致删除的结果不正确。
待删除的有序集合中,分数最好相同,否则删除结果会不正常。
 */
func Zremrangebylex(conn redis.Conn) {

}

/**
移除有序集key中，指定排名(rank)区间内的所有成员。下标参数start和stop都以0为底，0处是分数最小的那个元素。这些索引也可是负数，表示位移从最高分处开始数。例如，-1是分数最高的元素，-2是分数第二高的，依次类推。
 */
func Zremrangebyrank(conn redis.Conn) {

}

/**
移除有序集key中，所有score值介于min和max之间(包括等于min或max)的成员。 自版本2.1.6开始，score值等于min或max的成员也可以不包括在内，语法请参见ZRANGEBYSCORE命令。
 */
func Zremrangebyscore(conn redis.Conn) {

}

/**
返回有序集key中，指定区间内的成员。其中成员的位置按score值递减(从大到小)来排列。具有相同score值的成员按字典序的反序排列。 除了成员按score值递减的次序排列这一点外，ZREVRANGE命令的其他方面和ZRANGE命令一样。
 */
func Zrevrange(conn redis.Conn) {

}

/**
ZREVRANGEBYLEX 返回指定成员区间内的成员，按成员字典倒序排序, 分数必须相同。
在某些业务场景中,需要对一个字符串数组按名称的字典顺序进行倒序排列时,可以使用Redis中SortSet这种数据结构来处理。
 */
func Zrevrangebylex(conn redis.Conn) {

}

/**
ZREVRANGEBYSCORE 返回有序集合中指定分数区间内的成员，分数由高到低排序。
 */
func Zrevrangebyscore(conn redis.Conn) {

}

/**
返回有序集key中成员member的排名，其中有序集成员按score值从大到小排列。排名以0为底，也就是说，score值最大的成员排名为0。

使用ZRANK命令可以获得成员按score值递增(从小到大)排列的排名。
 */
func Zrevrank(conn redis.Conn) {

}

/**

 */
func Zscan(conn redis.Conn) {

}

/**
返回有序集key中，成员member的score值。

如果member元素不是有序集key的成员，或key不存在，返回nil。
 */
func Zscore(conn redis.Conn) {

}

/**

 */
func Zunionstore(conn redis.Conn) {

}