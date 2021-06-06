//-----------------------------------------------------------------------------
/*

Solids built with Bezier Curves

*/
//-----------------------------------------------------------------------------

package main

import (
	"log"

	"github.com/jakoblorz/sdfx/render"
	"github.com/jakoblorz/sdfx/sdf"
)

//-----------------------------------------------------------------------------

func bowlingPin() error {
	b := sdf.NewBezier()
	b.Add(0, 0)
	b.Add(2.031/2.0, 0).HandleFwd(sdf.DtoR(45), 2)
	b.Add(4.766/2.0, 4.5).Handle(sdf.DtoR(90), 2, 2)
	b.Add(1.797/2.0, 10).Handle(sdf.DtoR(90), 3, 3)
	b.Add(2.547/2.0, 13.5).Handle(sdf.DtoR(90), 1, 1)
	b.Add(0, 15).HandleRev(sdf.DtoR(0), 1)
	b.Close()

	p, err := b.Polygon()
	if err != nil {
		return err
	}
	s0, err := sdf.Polygon2D(p.Vertices())
	if err != nil {
		return err
	}
	s1, err := sdf.Revolve3D(s0)
	if err != nil {
		return err
	}
	render.RenderSTL(s1, 300, "bowlingpin.stl")
	return nil
}

//-----------------------------------------------------------------------------

func egg1() error {
	b := sdf.NewBezier()
	b.Add(0, 0).HandleFwd(sdf.DtoR(0), 10)
	b.Add(0, 16).HandleRev(sdf.DtoR(0), 5)
	b.Close()

	p, err := b.Polygon()
	if err != nil {
		return err
	}
	s0, err := sdf.Polygon2D(p.Vertices())
	if err != nil {
		return err
	}
	s1, err := sdf.Revolve3D(s0)
	if err != nil {
		return err
	}
	render.RenderSTL(s1, 300, "egg1.stl")
	return nil
}

//-----------------------------------------------------------------------------

func egg2() error {
	h := 8.0
	r := 2.5

	b := sdf.NewBezier()
	b.Add(0, 0).HandleFwd(sdf.DtoR(0), r/2)
	b.Add(r, 0.4*h).Handle(sdf.DtoR(90), 0.7*r, 0.7*r)
	b.Add(0, h).HandleRev(sdf.DtoR(0), r/3)
	b.Close()

	p, err := b.Polygon()
	if err != nil {
		return err
	}
	s0, err := sdf.Polygon2D(p.Vertices())
	if err != nil {
		return err
	}
	s1, err := sdf.Revolve3D(s0)
	if err != nil {
		return err
	}
	render.RenderSTL(s1, 300, "egg2.stl")
	return nil
}

//-----------------------------------------------------------------------------

func bowl() error {
	b := sdf.NewBezier()
	b.Add(1.428570, 0.000000)
	b.Add(194.311790, 1.616000).Mid()
	b.Add(424.623890, -2.388090).Mid()
	b.Add(584.285710, 98.571430)
	b.Add(730.711690, 191.161470).Mid()
	b.Add(845.816870, 372.034250).Mid()
	b.Add(850.576860, 545.337880)
	b.Add(855.789270, 735.113020).Mid()
	b.Add(679.478190, 877.171053).Mid()
	b.Add(586.270181, 1049.835600)
	b.Add(562.176808, 1094.467720).Mid()
	b.Add(551.662561, 1169.752600).Mid()
	b.Add(530.555895, 1191.428570)
	b.Add(506.830592, 1215.793810).Mid()
	b.Add(461.351740, 1202.110070).Mid()
	b.Add(444.285710, 1178.571430)
	b.Add(414.233090, 1137.120790).Mid()
	b.Add(452.788480, 1075.361930).Mid()
	b.Add(470.000000, 1027.142850)
	b.Add(531.988775, 853.477662).Mid()
	b.Add(743.353570, 724.365420).Mid()
	b.Add(743.662950, 546.384650)
	b.Add(743.899310, 410.411750).Mid()
	b.Add(648.722298, 272.112130).Mid()
	b.Add(536.903859, 194.747260)
	b.Add(387.543410, 91.407850).Mid()
	b.Add(0.000000, 101.890120).Mid()
	b.Add(0.000000, 101.890120)
	b.Close()

	p, err := b.Polygon()
	if err != nil {
		return err
	}
	s0, err := sdf.Polygon2D(p.Vertices())
	if err != nil {
		return err
	}
	s1, err := sdf.RevolveTheta3D(s0, sdf.DtoR(270))
	if err != nil {
		return err
	}
	render.RenderSTL(s1, 300, "bowl.stl")
	return nil
}

//-----------------------------------------------------------------------------

func vase() error {
	b := sdf.NewBezier()
	b.Add(0.357140, 0.000000)
	b.Add(286.591374, 0.000000)
	b.Add(286.591374, 0.000000).Mid()
	b.Add(438.180430, 405.704940).Mid()
	b.Add(438.180430, 599.667510)
	b.Add(438.180430, 747.833392).Mid()
	b.Add(340.551300, 912.588850).Mid()
	b.Add(326.785720, 937.142860)
	b.Add(313.020136, 961.696870).Mid()
	b.Add(289.271359, 1007.535800).Mid()
	b.Add(283.928573, 1025.714290)
	b.Add(278.585786, 1043.892790).Mid()
	b.Add(279.638315, 1091.954510).Mid()
	b.Add(254.642859, 1045.714290)
	b.Add(229.647403, 999.474070).Mid()
	b.Add(259.847631, 950.274470).Mid()
	b.Add(291.785716, 902.142860)
	b.Add(348.012060, 817.408149).Mid()
	b.Add(405.357150, 762.993097).Mid()
	b.Add(405.357150, 667.857140)
	b.Add(405.357150, 510.752570).Mid()
	b.Add(371.256950, 100.000000).Mid()
	b.Add(208.928573, 100.000000)
	b.Add(139.284800, 100.000000).Mid()
	b.Add(0.000000, 98.928570).Mid()
	b.Add(0.000000, 98.928570)
	b.Close()

	p, err := b.Polygon()
	if err != nil {
		return err
	}
	s0, err := sdf.Polygon2D(p.Vertices())
	if err != nil {
		return err
	}
	s1, err := sdf.Revolve3D(s0)
	if err != nil {
		return err
	}
	render.RenderSTL(s1, 300, "vase.stl")
	return nil
}

//-----------------------------------------------------------------------------

func shape() error {
	b := sdf.NewBezier()
	b.Add(-788.571430, 666.647920)
	b.Add(-788.785400, 813.701340).Mid()
	b.Add(-759.449240, 1023.568700).Mid()
	b.Add(-588.571430, 1026.647900)
	b.Add(-417.693610, 1029.727200).Mid()
	b.Add(-583.793160, 507.272270).Mid()
	b.Add(-285.714290, 506.647920)
	b.Add(12.364584, 506.023560).Mid()
	b.Add(-137.634380, 1110.386900).Mid()
	b.Add(85.714281, 1115.219300)
	b.Add(309.062940, 1120.051800).Mid()
	b.Add(498.298980, 1086.587000).Mid()
	b.Add(491.428570, 903.790780)
	b.Add(484.558160, 720.994550).Mid()
	b.Add(79.128329, 547.886390).Mid()
	b.Add(62.857140, 292.362210)
	b.Add(46.585951, 36.838026).Mid()
	b.Add(367.678530, -5.375978).Mid()
	b.Add(374.285720, -179.066370)
	b.Add(380.892900, -352.756760).Mid()
	b.Add(273.020040, -521.481290).Mid()
	b.Add(131.428570, -521.923510)
	b.Add(-10.162890, -522.365730).Mid()
	b.Add(50.355420, -54.901413).Mid()
	b.Add(-134.285720, -59.066363)
	b.Add(-318.926860, -63.231312).Mid()
	b.Add(-304.285720, -429.542560).Mid()
	b.Add(-442.857150, -433.352080)
	b.Add(-581.428570, -437.161610).Mid()
	b.Add(-750.919960, -371.353320).Mid()
	b.Add(-748.571430, -221.923510)
	b.Add(-746.222890, -72.493698).Mid()
	b.Add(-413.586510, -77.312515).Mid()
	b.Add(-402.857140, 120.933630)
	b.Add(-424.396820, 260.368600).Mid()
	b.Add(-788.357460, 519.594510).Mid()
	b.Add(-788.571430, 666.647920)
	b.Close()
	p, err := b.Polygon()
	if err != nil {
		return err
	}
	s0, err := sdf.Polygon2D(p.Vertices())
	if err != nil {
		return err
	}

	b = sdf.NewBezier()
	b.Add(37.142857, 663.790780)
	b.Add(-26.199008, 711.618410).Mid()
	b.Add(-3.917606, 881.472120).Mid()
	b.Add(60.000000, 943.790780)
	b.Add(123.917610, 1006.109400).Mid()
	b.Add(266.231230, 1026.016500).Mid()
	b.Add(314.285710, 958.076490)
	b.Add(362.340190, 890.136510).Mid()
	b.Add(272.791650, 783.239950).Mid()
	b.Add(220.000000, 732.362200)
	b.Add(167.208350, 681.484450).Mid()
	b.Add(100.484720, 615.963150).Mid()
	b.Add(37.142857, 663.790780)
	b.Close()
	p, err = b.Polygon()
	if err != nil {
		return err
	}
	s1, err := sdf.Polygon2D(p.Vertices())
	if err != nil {
		return err
	}

	b = sdf.NewBezier()
	b.Add(105.714290, -381.923510)
	b.Add(72.682777, -314.172080).Mid()
	b.Add(-4.286470, -63.657320).Mid()
	b.Add(57.142857, -1.923510)
	b.Add(118.572180, 59.810298).Mid()
	b.Add(205.984530, -33.044213).Mid()
	b.Add(237.142850, -110.494940)
	b.Add(268.301180, -187.945670).Mid()
	b.Add(299.753500, -280.841720).Mid()
	b.Add(262.857150, -350.494940)
	b.Add(233.564570, -405.793690).Mid()
	b.Add(138.745800, -449.674940).Mid()
	b.Add(105.714290, -381.923510)
	b.Close()
	p, err = b.Polygon()
	if err != nil {
		return err
	}
	s2, err := sdf.Polygon2D(p.Vertices())
	if err != nil {
		return err
	}

	b = sdf.NewBezier()
	b.Add(-668.571430, -247.637800)
	b.Add(-682.865090, -172.496010).Mid()
	b.Add(-553.303000, -108.477610).Mid()
	b.Add(-491.428580, -96.209224)
	b.Add(-429.554160, -83.940837).Mid()
	b.Add(-373.722060, -101.371460).Mid()
	b.Add(-351.428570, -159.066370)
	b.Add(-329.135080, -216.761280).Mid()
	b.Add(-375.302040, -306.770950).Mid()
	b.Add(-445.714290, -333.352080)
	b.Add(-516.126540, -359.933210).Mid()
	b.Add(-654.277770, -322.779580).Mid()
	b.Add(-668.571430, -247.637800)
	b.Close()
	p, err = b.Polygon()
	if err != nil {
		return err
	}
	s3, err := sdf.Polygon2D(p.Vertices())
	if err != nil {
		return err
	}

	b = sdf.NewBezier()
	b.Add(-697.142850, 569.505060)
	b.Add(-746.320070, 637.664520).Mid()
	b.Add(-737.208530, 704.563220).Mid()
	b.Add(-702.857140, 752.362200)
	b.Add(-670.958410, 796.748390).Mid()
	b.Add(-623.040800, 829.292720).Mid()
	b.Add(-568.571430, 763.790770)
	b.Add(-514.102060, 698.288820).Mid()
	b.Add(-455.211270, 526.871080).Mid()
	b.Add(-520.000000, 500.933640)
	b.Add(-561.944210, 484.141750).Mid()
	b.Add(-647.965630, 501.345600).Mid()
	b.Add(-697.142850, 569.505060)
	b.Close()
	p, err = b.Polygon()
	if err != nil {
		return err
	}
	s4, err := sdf.Polygon2D(p.Vertices())
	if err != nil {
		return err
	}

	b = sdf.NewBezier()
	b.Add(-288.571430, 106.647920)
	b.Add(-337.549290, 162.917530).Mid()
	b.Add(-355.864180, 274.511310).Mid()
	b.Add(-302.857140, 338.076490)
	b.Add(-249.850100, 401.641670).Mid()
	b.Add(-107.428130, 437.061600).Mid()
	b.Add(-40.000000, 369.505060)
	b.Add(27.428134, 301.948520).Mid()
	b.Add(-6.357993, 114.855610).Mid()
	b.Add(-77.142857, 63.790776)
	b.Add(-147.927720, 12.725947).Mid()
	b.Add(-239.593570, 50.378315).Mid()
	b.Add(-288.571430, 106.647920)
	b.Close()
	p, err = b.Polygon()
	if err != nil {
		return err
	}
	s5, err := sdf.Polygon2D(p.Vertices())
	if err != nil {
		return err
	}

	s2d := sdf.Difference2D(s0, sdf.Union2D(s1, s2, s3, s4, s5))
	s3d, err := sdf.ExtrudeRounded3D(s2d, 200, 20)
	if err != nil {
		return err
	}
	render.RenderSTL(s3d, 300, "shape.stl")
	return nil
}

//-----------------------------------------------------------------------------

func main() {
	err := bowlingPin()
	if err != nil {
		log.Fatalf("error: %s\n", err)
	}
	err = bowl()
	if err != nil {
		log.Fatalf("error: %s\n", err)
	}
	err = vase()
	if err != nil {
		log.Fatalf("error: %s\n", err)
	}
	err = egg1()
	if err != nil {
		log.Fatalf("error: %s\n", err)
	}
	err = egg2()
	if err != nil {
		log.Fatalf("error: %s\n", err)
	}
	err = shape()
	if err != nil {
		log.Fatalf("error: %s\n", err)
	}
}

//-----------------------------------------------------------------------------
