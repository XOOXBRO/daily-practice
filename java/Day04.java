/**
 * Created by 王盈 on 2023/1/13.
 * 给你两个下标从 0 开始的字符串 s 和 target 。你可以从 s 取出一些字符并将其重排，得到若干新的字符串。
 从 s 中取出字符并重新排列，返回可以形成 target 的 最大 副本数。
 */
public class Day04 {
    public int rearrangeCharacters(String s, String target) {
        int sum=0;
        int index=0;
        for(int i=0;i<target.length();i++){
            char c = target.charAt(i);
            int i1 = s.indexOf(c, index);
            index=i1+1;
            if(i1==-1){
                break;
            }
            if(i==target.length()-1){
                sum++;
                i=-1;
            }
        }
        System.out.println(sum);
        return sum;
    }

    public static void main(String[] args) {
        Day04 day04=new Day04();
        day04.rearrangeCharacters("ilovecodingonleetcode","code");
    }
}
