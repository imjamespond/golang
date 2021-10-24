import com.sun.jna.Library;
import com.sun.jna.Memory;
import com.sun.jna.Native;
import com.sun.jna.Pointer;

import javax.security.auth.callback.Callback;
import java.nio.charset.StandardCharsets;

interface JNA extends Library {
    JNA instance = (JNA )  Native.load("test", JNA.class);
    public interface Func extends Callback {
        int invoke(int a,int b);
    }
    int test(Func c);
    int add(int a, int b);
    Pointer hello(Pointer name);
    void freePoint(Pointer pointer);
}

public class JNATest{
    public static void main(String[] args) {
        {
            int result= JNA.instance.test(new JNA.Func() {
                @Override
                public int invoke(int a, int b) {
                    System.out.println(a * b);
                    return a*b;
                }
            });
            System.out.println(result);
        }


        {
           int rs = JNA.instance.add(1,2);
           System.out.println(rs);
        }

        {
            String name = "foobar";
            // 申请jna的内存空间
            Pointer pname = new Memory(name.getBytes(StandardCharsets.UTF_8).length + 1);
            // 设置传入参数值
            pname.setString(0, name);
            Pointer ptr = null;
            try {
                ptr = JNA.instance.hello(pname);
                System.out.println(ptr.getString(0, "utf8"));
            } finally {
                // 释放传入jna的指针对应的内存空间
                Native.free(Pointer.nativeValue(pname));
                // 解决多次调用崩溃的问题
                Pointer.nativeValue(pname, 0);
                if (ptr != null) {
                    // 释放go中申请的C的内存
                    JNA.instance.freePoint(ptr);
                }
            }
        }
    }
}

