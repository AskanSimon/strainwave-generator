package main

import (
    "bytes"
    "fmt"
    "log"
    "os"

    "github.com/ctessum/go.clipper"
)

const scaleFactor = 1e9

func scaleUp(value float64) clipper.CInt {
    return clipper.CInt(value * scaleFactor)
}

func scaleDown(value clipper.CInt) float64 {
    return float64(value) / scaleFactor
}

func writeSvg(data []byte) error {
    f, err := os.OpenFile("test.svg", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
    if err != nil {
        return err
    }
    defer f.Close()

    _, err = f.Write(data)
    if err != nil {
        return err
    }

    return nil
}

// 64bit Zahl +- 9223372036854775807 // 19 stellen

func main() {
    clip := clipper.NewClipper(clipper.IoNone)
    pclip := clipper.Paths{{{1440000000000, 900000000000}, {1439666234781, 921458904966}, {1438665262022, 942897049827}, {1436998050098, 964293694564}, {1434666211930, 985628139303}, {1431672003421, 1006879744347}, {1428018321273, 1028027950142}, {1423708700188, 1049052297163}, {1418747309447, 1069932445714}, {1413138948875, 1090648195598}, {1406889044197, 1111179505665}, {1400003641794, 1131506513200}, {1392489402847, 1151609553135}, {1384353596897, 1171469177077}, {1375604094812, 1191066172122}, {1366249361168, 1210381579445}, {1356298446068, 1229396712635}, {1345760976378, 1248093175779}, {1334647146420, 1266452881257}, {1322967708107, 1284458067239}, {1310733960541, 1302091314871}, {1297957739082, 1319335565125}, {1284651403900, 1336174135302}, {1270827828015, 1352590735172}, {1256500384844, 1368569482737}, {1241682935261, 1384094919589}, {1226389814194, 1399152025871}, {1210635816749, 1413726234805}, {1194436183900, 1427803446786}, {1177806587747, 1441370043019}, {1160763116350, 1454412898700}, {1143322258167, 1466919395708}, {1125500886101, 1478877434814}, {1107316241179, 1490275447388}, {1088785915871, 1501102406588}, {1069927837069, 1511347838033}, {1050760248748, 1521001829927}, {1031301694309, 1530055042657}, {1011570998649, 1538498717825}, {991587249942, 1546324686721}, {971369781172, 1553525378224}, {950938151436, 1560093826131}, {930312127016, 1566023675893}, {909511662259, 1571309190763}, {888556880270, 1575945257346}, {867468053448, 1579927390549}, {846265583871, 1583251737913}, {824969983558, 1585915083348}, {803601854629, 1587914850239}, {782181869367, 1589249103941}, {760730750225, 1589916553649}, {739269249774, 1589916553649}, {717818130632, 1589249103941}, {696398145370, 1587914850239}, {675030016441, 1585915083348}, {653734416128, 1583251737913}, {632531946551, 1579927390549}, {611443119729, 1575945257346}, {590488337740, 1571309190763}, {569687872983, 1566023675893}, {549061848563, 1560093826131}, {528630218827, 1553525378224}, {508412750057, 1546324686721}, {488429001350, 1538498717825}, {468698305690, 1530055042657}, {449239751251, 1521001829927}, {430072162930, 1511347838033}, {411214084128, 1501102406588}, {392683758820, 1490275447388}, {374499113898, 1478877434814}, {356677741832, 1466919395708}, {339236883649, 1454412898700}, {322193412252, 1441370043019}, {305563816099, 1427803446786}, {289364183250, 1413726234805}, {273610185805, 1399152025871}, {258317064738, 1384094919589}, {243499615155, 1368569482737}, {229172171984, 1352590735172}, {215348596099, 1336174135302}, {202042260917, 1319335565125}, {189266039458, 1302091314871}, {177032291892, 1284458067239}, {165352853579, 1266452881257}, {154239023621, 1248093175779}, {143701553931, 1229396712635}, {133750638831, 1210381579445}, {124395905187, 1191066172122}, {115646403102, 1171469177077}, {107510597152, 1151609553135}, {99996358205, 1131506513200}, {93110955802, 1111179505665}, {86861051124, 1090648195598}, {81252690552, 1069932445714}, {76291299811, 1049052297163}, {71981678726, 1028027950142}, {68327996578, 1006879744347}, {65333788069, 985628139303}, {63001949901, 964293694564}, {61334737977, 942897049827}, {60333765218, 921458904966}, {60000000000, 900000000000}, {60333765218, 878541095033}, {61334737977, 857102950172}, {63001949901, 835706305435}, {65333788069, 814371860696}, {68327996578, 793120255652}, {71981678726, 771972049857}, {76291299811, 750947702836}, {81252690552, 730067554285}, {86861051124, 709351804401}, {93110955802, 688820494334}, {99996358205, 668493486799}, {107510597152, 648390446864}, {115646403102, 628530822922}, {124395905187, 608933827877}, {133750638831, 589618420554}, {143701553931, 570603287364}, {154239023621, 551906824220}, {165352853579, 533547118742}, {177032291892, 515541932760}, {189266039458, 497908685128}, {202042260917, 480664434874}, {215348596099, 463825864697}, {229172171984, 447409264827}, {243499615155, 431430517262}, {258317064738, 415905080410}, {273610185805, 400847974128}, {289364183250, 386273765194}, {305563816099, 372196553213}, {322193412252, 358629956980}, {339236883649, 345587101299}, {356677741832, 333080604291}, {374499113898, 321122565185}, {392683758820, 309724552611}, {411214084128, 298897593411}, {430072162930, 288652161966}, {449239751251, 278998170072}, {468698305690, 269944957342}, {488429001350, 261501282174}, {508412750057, 253675313278}, {528630218827, 246474621775}, {549061848563, 239906173868}, {569687872983, 233976324106}, {590488337740, 228690809236}, {611443119729, 224054742653}, {632531946551, 220072609450}, {653734416128, 216748262086}, {675030016441, 214084916651}, {696398145370, 212085149760}, {717818130632, 210750896058}, {739269249774, 210083446350}, {760730750225, 210083446350}, {782181869367, 210750896058}, {803601854629, 212085149760}, {824969983558, 214084916651}, {846265583871, 216748262086}, {867468053448, 220072609450}, {888556880270, 224054742653}, {909511662259, 228690809236}, {930312127016, 233976324106}, {950938151436, 239906173868}, {971369781172, 246474621775}, {991587249942, 253675313278}, {1011570998649, 261501282174}, {1031301694309, 269944957342}, {1050760248748, 278998170072}, {1069927837069, 288652161966}, {1088785915871, 298897593411}, {1107316241179, 309724552611}, {1125500886101, 321122565185}, {1143322258167, 333080604291}, {1160763116350, 345587101299}, {1177806587747, 358629956980}, {1194436183900, 372196553213}, {1210635816749, 386273765194}, {1226389814194, 400847974128}, {1241682935261, 415905080410}, {1256500384844, 431430517262}, {1270827828015, 447409264827}, {1284651403900, 463825864697}, {1297957739082, 480664434874}, {1310733960541, 497908685128}, {1322967708107, 515541932760}, {1334647146420, 533547118742}, {1345760976378, 551906824220}, {1356298446068, 570603287364}, {1366249361168, 589618420554}, {1375604094812, 608933827877}, {1384353596897, 628530822922}, {1392489402847, 648390446864}, {1400003641794, 668493486799}, {1406889044197, 688820494334}, {1413138948875, 709351804401}, {1418747309447, 730067554285}, {1423708700188, 750947702836}, {1428018321273, 771972049857}, {1431672003421, 793120255652}, {1434666211930, 814371860696}, {1436998050098, 835706305436}, {1438665262022, 857102950172}, {1439666234781, 878541095033}}}

    clip.AddPaths(psubj, clipper.PtSubject, true)
    clip.AddPaths(pclip, clipper.PtClip, true)
    // CtUnion (auch viele)
    // CtIntersection
    // CtDifference
    combinedPaths, ok := clip.Execute1(clipper.CtUnion, clipper.PftNonZero, clipper.PftNonZero) // ???
    if !ok {
        panic("failed to execute")
    }
    /*
        if len(combinedPaths) > 1 {
            panic("unexpected multiple paths in combined paths")
        }
    */

    var b bytes.Buffer
    b.WriteString("<svg width=\"1500\" height=\"2000\" xmlns=\"http://www.w3.org/2000/svg\">")
    b.WriteString("<g stroke=\"black\" stroke-width=\"0.5\" fill=\"none\">")

    for _, path := range combinedPaths {
        b.WriteString("<path d=\"M")
        first := true
        for _, point := range path {

            b.WriteString(fmt.Sprintf("%.14f %.14f ", scaleDown(point.X), scaleDown(point.Y)))
            if first {
                b.WriteString("L ")
                first = false
            }
        }
        b.WriteString("z\"/>\n")
    }

    b.WriteString("</g></svg>")

    err := writeSvg(b.Bytes())
    if err != nil {
        log.Fatal(err)
    }
}