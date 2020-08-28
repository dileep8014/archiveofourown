import { message } from 'antd';

const numberParseChina = function(number: number) {

  //汉字的数字
  const cnNums = ['零', '一', '二', '三', '四', '五', '六', '七', '八', '九'];
  //基本单位
  const cnIntRadice = ['', '十', '百', '千'];
  //最大处理的数字
  const maxNum = 1000;
  //输出的中文金额字符串
  let chineseStr = '';
  if (number >= maxNum) {
    message.error('到达最大分卷数');
    return;
  }

  //转换为字符串
  let integerNum = number.toString();
  //获取整型部分转换
  if (parseInt(integerNum, 10) > 0) {
    let zeroCount = 0;
    let IntLen = integerNum.length;
    for (let i = 0; i < IntLen; i++) {
      // var n = integerNum.substr(i, 1);
      let n = integerNum.substring(i, i + 1);
      let p = IntLen - i - 1;
      let q = p / 4;
      let m = p % 4;
      if (n == '0') {
        zeroCount++;
      } else {
        if (zeroCount > 0) {
          chineseStr += cnNums[0];
        }
        //归零
        zeroCount = 0;
        chineseStr += cnNums[parseInt(n)] + cnIntRadice[m];
      }
    }
  }
  if (chineseStr == '') {
    chineseStr += cnNums[0];
  }

  return chineseStr;
};

export default numberParseChina;
