set terminal postscript eps monochrome size 3,2.1
set output "pdf.ps"
set xlabel "Shustring Length"
set ylabel "P"
set arrow nohead from 10,0 to 10,0.5 ls 3
set arrow nohead from 11.38,0 to 11.38,0.5 ls 3
set arrow nohead from 13,0 to 13,0.5 ls 3
plot[8:17][*:*] "-" t "exp" w l, "-" t "obs" w l
1	0.000000
2	0.000000
3	0.000000
4	0.000000
5	0.000000
6	0.000000
7	0.000000
8	0.000000
9	0.000486
10	0.147987
11	0.472270
12	0.266878
13	0.083015
14	0.021940
15	0.005562
16	0.001395
17	0.000349
18	0.000087
19	0.000022
20	0.000005
21	0.000001
22	0.000000
23	0.000000
24	0.000000
25	0.000000
26	0.000000
27	0.000000
e
8	0.000000
9	0.000480005
10	0.147499
11	0.472788
12	0.26678
13	0.0831958
14	0.0219452
15	0.00547205
16	0.00138001
17	0.000342003
18	7.90008e-05
19	2.30002e-05
20	5.00005e-06
21	5.00005e-06
22	2.00002e-06
23	2.00002e-06
24	2.00002e-06