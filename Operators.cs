


//public static DateTime operator +(DateTime d, TimeSpan t);
//public static TimeSpan operator -(DateTime d1, DateTime d2);
//public static DateTime operator -(DateTime d, TimeSpan t);
//public static bool operator ==(DateTime d1, DateTime d2);
//public static bool operator !=(DateTime d1, DateTime d2);
//public static bool operator <(DateTime t1, DateTime t2);
//public static bool operator >(DateTime t1, DateTime t2);
//public static bool operator <=(DateTime t1, DateTime t2);
//public static bool operator >=(DateTime t1, DateTime t2);

//public class Range : IReadOnlyList<int>
//{
//    public int Start { get; private set; }
//    public int Count { get; private set; }
//    public int this[int index]
//    {
//        get
//        {
//            if (index < 0 || index >= Count)
//            {
//                throw new IndexOutOfBoundsException("index");
//            }
//            return Start + index;
//        }
//    }
//    public Range(int start, int count)
//    {
//        this.Start = start;
//        this.Count = count;
//    }
//    public IEnumerable<int> GetEnumerator()
//    {
//        return Enumerable.Range(Start, Count);
//    }
//}