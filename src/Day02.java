/**
 * Created by 王盈 on 2023/1/11.
 * 给你一个下标从 0 开始长度为 n 的字符串 num ，它只包含数字。
 如果对于 每个 0 <= i < n 的下标 i ，都满足数位 i 在 num 中出现了 num[i]次，那么请你返回 true ，否则返回 false 。
 */
public class Day02 {
    public boolean digitCount(String num) {
        for(int i=0;i<num.length();i++){
            int sum=0;
            for(int j=0;j<num.length();j++){
                if(String.valueOf(i).equals(String.valueOf(num.charAt(j)))){
                    sum++;
                }
            }
            if(Integer.parseInt(String.valueOf(num.charAt(i)))!=sum){
                return false;
            }
        }
        return true;
    }

    public static void main(String[] args) {
        Day02 day02=new Day02();
        day02.digitCount("1210");
    }
}
