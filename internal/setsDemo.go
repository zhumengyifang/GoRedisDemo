package internal

import "github.com/garyburd/redigo/redis"

/**
添加一个或多个指定的member元素到集合的 key中.指定的一个或者多个元素member
如果已经在集合key中存在则忽略.如果集合key 不存在，则新建集合key,并添加member元素到集合key中.
*/
func Sadd(conn redis.Conn) {

}

/**
返回集合存储的key的基数 (集合元素的数量)
*/
func Scard(conn redis.Conn) {

}

/**
返回一个集合与给定集合的差集的元素.
*/
func Sdiff(conn redis.Conn) {

}

/**
该命令类似于 SDIFF, 不同之处在于该命令不返回结果集，而是将结果存放在destination集合中.
如果destination已经存在, 则将其覆盖重写.
*/
func SdiffStore(conn redis.Conn) {

}

/**
返回指定所有的集合的成员的交集.
*/
func Sinter(conn redis.Conn) {

}

/**
这个命令与SINTER命令类似, 但是它并不是直接返回结果集,而是将结果保存在 destination集合中.
如果destination 集合存在, 则会被重写.
*/
func SinterStore(conn redis.Conn) {

}

/**
返回成员 member 是否是存储的集合 key的成员.
*/
func SisMember(conn redis.Conn) {

}

/**
返回key集合所有的元素.
该命令的作用与使用一个参数的SINTER 命令作用相同.
*/
func Smembers(conn redis.Conn) {

}

/**
将member从source集合移动到destination集合中. 对于其他的客户端,在特定的时间元素将会作为source或者destination集合的成员出现.

如果source 集合不存在或者不包含指定的元素,这smove命令不执行任何操作并且返回0.否则对象将会从source集合中移除，并添加到destination集合中去，如果destination集合已经存在该元素，则smove命令仅将该元素充source集合中移除. 如果source 和destination不是集合类型,则返回错误.
*/
func Smove(conn redis.Conn) {

}

/**
从存储在key的集合中移除并返回一个或多个随机元素。

此操作与SRANDMEMBER类似，它从一个集合中返回一个或多个随机元素，但不删除元素。
*/
func Spop(conn redis.Conn) {

}

/**
仅提供key参数，那么随机返回key集合中的一个元素.

Redis 2.6开始，可以接受 count 参数，如果count是整数且小于元素的个数，返回含有 count 个不同的元素的数组，如果count是个整数且大于集合中元素的个数时，仅返回整个集合的所有元素，当count是负数，则会返回一个包含count的绝对值的个数元素的数组，如果count的绝对值大于元素的个数，则返回的结果集里会出现一个元素出现多次的情况.

仅提供key参数时，该命令作用类似于SPOP命令，不同的是SPOP命令会将被选择的随机元素从集合中移除，而SRANDMEMBER仅仅是返回该随记元素，而不做任何操作.
*/
func SrandMember(conn redis.Conn) {

}

/**
在key集合中移除指定的元素. 如果指定的元素不是key集合中的元素则忽略 如果key集合不存在则被视为一个空的集合，该命令返回0.

如果key的类型不是一个集合,则返回错误.
*/
func Sren(conn redis.Conn) {

}

/**

 */
func Sscan(conn redis.Conn) {

}

/**
返回给定的多个集合的并集中的所有成员.
*/
func Sunion(conn redis.Conn) {

}

/**
该命令作用类似于SUNION命令,不同的是它并不返回结果集,而是将结果存储在destination集合中.

如果destination 已经存在,则将其覆盖.
*/
func SunionStore(conn redis.Conn) {

}
