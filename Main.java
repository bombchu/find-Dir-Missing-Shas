package xyz.bombchu;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.*;
import java.util.regex.Pattern;
import java.util.stream.Collectors;
import java.util.stream.Stream;

public class Main {

    public static void main(String[] args) {

        String filePathString = "/home/bombchu/mnt-points/1/bak/master0/no-clobber/dta/";

        List<String> result0 = new LinkedList<>();
        List<String> result1 = new LinkedList<>();

        try (Stream<Path> walk = Files.walk(Paths.get(filePathString))) {

            result0 = walk.map(x -> x.toString()).filter(f -> f.endsWith(".mp4")).collect(Collectors.toList());

        } catch (IOException e) {
            e.printStackTrace();
        }

        try (Stream<Path> walk = Files.walk(Paths.get(filePathString))) {

            result1 = walk.map(x -> x.toString()).filter(f -> f.endsWith(".rar")).collect(Collectors.toList());

            /*for (String st : result1) {
                System.out.println(st);
            }
            System.out.println("end 1 sout");*/

            result0.addAll(result1);

            /*for (String s : result0) {
                System.out.println(s);
            }
            System.out.println("end 2 sout");*/

        } catch (IOException e) {
            e.printStackTrace();
        }

        TreeSet<String> finalList = null;

        try (Stream<Path> walk2 = Files.walk(Paths.get(filePathString))) {

            List<String> result2 = walk2.map(x -> x.toString()).filter(f -> f.endsWith(".sha256")).collect(Collectors.toList());

            /*for (String str : result2) {
                System.out.println(str);
            }
            System.out.println("end 3 sout");*/

            Pattern pattern = Pattern.compile("/.+/");

            finalList = result0.parallelStream().filter(item0 -> result2.parallelStream().noneMatch(item2 -> item2.contains(item0)))
                    .map(x -> new Scanner(x).findAll(pattern).map(m -> m.group()).collect(Collectors.toList()))
                    .flatMap(List::stream).collect(Collectors.toCollection(TreeSet::new));

        } catch (IOException e) {
            e.printStackTrace();
        }

        finalList.forEach(x -> System.out.println(x));
    }
}
