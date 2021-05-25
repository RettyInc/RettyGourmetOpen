const data = [
    // 店名, id, 昼予算, 夜予算, 行きたい数, 個室あり, カウンターあり, 喫煙・禁煙, PayPayが使える
    "いいおか 潮騒ホテルレストラン, 100001510561, 2000, 4000, 212, 1, 1, 1, 1",
    "うどん 丸香, 100000000614, 1000, 2500, 198, 0, 1, 0, 1",
    "かあさん 新宿駅前店, 100000093866, 1000, 3000, 265, 0, 1, 1, 1",
    "つじ半, 100000729115, 1000, 2000, 240, 0, 1, 0, 1",
    "あなご屋 銀座ひらい, 100000821612, 2000, 5000, 290, 1, 0, 1, 1",
    "おきなわ家, 100000949764, 1000, 3000, 232, 1, 1, 1, 0",
    "どい亭 ヒコボシ, 100000917872, 1500, 2500, 202, 1, 1, 0, 1",
    "ほんだ, 100000785336, 1500, 5000, 285, 1, 1, 1, 1",
  ]
  
  // csv行をカンマ区切りして、前後の空白を除去
  const csvArray = data.map(d => d.split(",").map(s => s.trim()))
  
  // paypay使えるを優先
  var canPaypayUse = []
  var cannotPaypayUse = []
  csvArray.forEach(function(csv) {
    if (csv[8] == 1) {
      canPaypayUse.push(csv)
    } else {
      cannotPaypayUse.push(csv)
    }
  })
  
  // 行きたい数でソート
  function sorter(a, b) {
    const aValue = parseInt(a[4], 10)
    const bValue = parseInt(b[4], 10)
    if (aValue < bValue) {
      return 1
    }
    if (aValue > bValue) {
      return -1
    }
    return 0
  }
  canPaypayUse.sort(sorter)
  cannotPaypayUse.sort(sorter)
  
  const sortedList = canPaypayUse.concat(cannotPaypayUse)
  
  // 先頭を取り出しお店IDを出力
  const resultRestaurant = sortedList[0]
  console.log(resultRestaurant[1])
  
  
  /*
  - 問題自体はデータ数少ないので手で溶ける問題
    - データ件数が多いときに同様に運用できるコードか
    - 入力データ形式が汎用的になっているか
  - 処理自体のパフォーマンス
  - 違うソート条件の指定が容易か
  
  */