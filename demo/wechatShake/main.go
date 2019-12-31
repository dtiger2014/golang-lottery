/**
 * 微信摇一摇
 * 增加互斥锁，保证并发更新数据的安全
 * 基础功能：
 * /lucky 只有一个抽奖的接口，奖品信息都是预先配置好的
 * 测试方法：
 * curl http://localhost:8080/
 * curl http://localhost:8080/lucky
 * 压力测试：（线程不安全的时候，总的中奖纪录会超过总的奖品数）
 * wrk -t10 -c10 -d5 http://localhost:8080/lucky
 */

package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
	"sync"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

const (
	giftTypeCoin = iota //虚拟币
	giftTypeCoupon // 优惠券，不相同的编码
	giftTypeCouponFix // 优惠券，相同的编码
	giftTypeRealSmall //
)