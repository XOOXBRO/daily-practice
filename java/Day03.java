import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * Created by 王盈 on 2023/1/12.
 * 给你一个字符串 s ，它包含一些括号对，每个括号中包含一个 非空 的键。
 比方说，字符串 "(name)is(age)yearsold" 中，有 两个 括号对，分别包含键 "name" 和 "age" 。
 你知道许多键对应的值，这些关系由二维字符串数组 knowledge 表示，其中 knowledge[i] = [keyi, valuei] ，表示键 keyi 对应的值为 valuei 。
 你需要替换 所有 的括号对。当你替换一个括号对，且它包含的键为 keyi 时，你需要：
 将 keyi 和括号用对应的值 valuei 替换。
 如果从 knowledge 中无法得知某个键对应的值，你需要将 keyi 和括号用问号 "?" 替换（不需要引号）。
 knowledge 中每个键最多只会出现一次。s 中不会有嵌套的括号。
 请你返回替换 所有 括号对后的结果字符串

 */
public class Day03 {
    public String evaluate(String s, List<List<String>> knowledge) {
        Map<String,String> map=new HashMap<String,String>();
        Map<String,String> replace=new HashMap<String,String>();
        knowledge.forEach(r->{
            if(r.size()==2){
                map.put(r.get(0),r.get(1));
            }
        });
        for(int i=0;i<s.length();i++){
            if(String.valueOf(s.charAt(i)).equals("(")){
                int i1 = s.indexOf(")", i + 1);
                String s1 = s.substring(i, i1 + 1);
                String s2 = s.substring(i + 1, i1);
                i=i1;
                String v1 = map.get(s2);
                if(v1==null){
                    replace.put(s1, "?");
                    continue;
                }
                replace.put(s1, v1);
            }
        }
        if(!replace.isEmpty()){
            for (String i : replace.keySet()) {
                s = s.replace(i, replace.get(i));
            }
            System.out.println(s);
            return s;
        }
        return s;
    }

    public static void main(String[] args) {
        Day03 day03=new Day03();
        List<List<String>> lists=new ArrayList<List<String>>();
        List<String> list1=new ArrayList<String>();
        List<String> list2=new ArrayList<String>();
        list1.add("a");
        list1.add("b");
        //list2.add("age");
        //list2.add("two");
        lists.add(list1);
        //lists.add(list2);
        day03.evaluate("(fy)(kj)(ege)r",lists);
    }
}
