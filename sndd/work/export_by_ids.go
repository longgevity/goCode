package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	// "regexp"
	"golang.org/x/net/html"
	"os"
	"strings"
)

// type User struct {
// 	Id   int
// 	Name string
// }

type Answer struct {
	Instruction string `json:"instruction"`
	Input       string `json:"input"`
	Output      string `json:"output"`
}

type Content struct {
	Content string `json:"content"`
}

// type Comment struct {
// 	Comment string `json:"content"`
// }

func main() {
	qids := [...]string{
		"758157094477238272",
		"759223811118862336",
		"762775257995153408",
		"763538577236824064",
		"766052542726672384",
		"767547618699120640",
		"770210015561781248",
		"770804559135248384",
		"772970615522070528",
		"773333521325494272",
		// "775115603735351296",
		"775917577808515072",
		"775901721783308288",
		"778766899331862528",
		"783075448606298112",
		"786029583395852288",
		"787084097330745344",
		"790688156738523136",
		"793211460191916032",
		"793979697007562752",
		// "795567245815320576",
		"759540771379613696",
		"762041892643934208",
		"764604649780154368",
		"767267007606427648",
		"769712558792904704",
		"771945456413577216",
		"774742775626338304",
		"777261062704730112",
		"787155194931712000",
		"792479980285202432",
		"792582096458747904",
		"759567637909999616",
		"762445955462205440",
		"764997817591795712",
		"767501548724031488",
		"770069259513630720",
		"775058564413788160",
		"777331264318476288",
		"785275354389549056",
		"787795006281879552",
		"790372333494865920",
		"792175818989244416",
		"795098975173742592",
		"797248713633435648",
		// "763512490003927040",
		// "766014990619643904",
		// "771151863931342848",
		"772590216287490048",
		"773673098510733312",
		"775250573111988224",
		"775875953007136768",
		"778750461468610560",
		"783582883855077376",
		"790462545797451776",
		// "792663176192856064",
		"793631214148915200",
		"795203635846975488",
		// "759607707677364224",
		"764723855494352896",
		"767180341134888960",
		// "771904440348315648",
		"773377891332395008",
		"776232420029632512",
		// "778536614375854080",
		"783460303152222208",
		// "784946540237164544",
		"786385231618904064",
		// "787086291736072192",
		"790755485358886912",
		"792171842206044160",
		"758096169875410944",
		// "763939267943731200",
		"766075731863277568",
		"768267103525539840",
		"769830603699916800",
		"773681471356538880",
		"778013421051449344",
		"786406113879920640",
		"786742498960740352",
		"788548430959087616",
		"791453305493655552",
		"768993291051077632",
		"772618651558875136",
		"775131145330364416",
		"777379138225967104",
		"790058054006738944",
		"795430887666880512",
	}

	len_qids := len(qids)

	// export string
	for index, qid := range qids {
		// fmt.Println(index)
		result, ans_len := export_by_ids(qid)

		file, err := os.Create("results/" + qid + ".json")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer file.Close()

		_, err = file.WriteString(result)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("%d/%d: %s.json (%d answers)...\n", index+1, len_qids, qid, ans_len)
	}

	// // export list
	// var answers []Answer
	// for index, qid := range qids {
	// 	answers_ := export_list_by_ids(qid)
	// 	answers = append(answers, answers_...)
	// 	fmt.Printf("%d/%d: %s...\n", index+1, len_qids, qid)
	// }

	// // 将结构体数组转为JSON格式并输出到控制台
	// jsonBytes, err := json.Marshal(answers)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// file, err2 := os.Create("test.json")
	// if err2 != nil {
	// 	fmt.Println("Error:", err2)
	// 	return
	// }
	// defer file.Close()

	// _, err = file.WriteString(string(jsonBytes))
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
}

func export_by_ids(qid string) (string, int) {

	sql1 := fmt.Sprintf(`
		SELECT yw.content instruction, ycs.comment_content input,yc.comment_content output,ycrs.mark_num mark_num
		 FROM yuwen_comment_relation ycrs
	     left join yuwen_comment ycs on ycs.comment_relation_pid = ycrs.rela_pid
	     left join yuwen_comment_relation ycr on ycr.parent_id = ycrs.rela_pid
	     left join yuwen_comment yc on yc.comment_relation_pid = ycr.rela_pid
	     left join yuwen_schedule_work ysw on ysw.id = ycr.schedule_work_id
	     left join yuwen_work yw on ysw.work_id = yw.id
	     where  yw.id='%s' and ycs.comment_content not like "%%\":\"https\:/%%"
	    --  where  yw.content like "%%请用自然观察法和生活观察法观察你喜欢的小动物%%"
	    --  or  yw.content like "%%训练重点：用想象力魔法棒“变环境”，打开想象编写故事%%"
	    --  and ycs.comment_reply_time < "2023-01-01 00:00:00"
	     and ycrs.identity =1 and ycrs.comment_type = 1 and yw.question_type =0
	    --  and ycrs.mark_num = 5
	    ORDER BY ycs.comment_reply_time
	    -- ORDER BY RAND()
	    limit 3000
		`, qid)
	// fmt.Println(sql1)

	// 连接数据库
	db, err := sql.Open("mysql", "yuwen_read:3t*hnCqIqTaE3woP@tcp(rr-2ze5mv4l0fp4c0dn3.mysql.rds.aliyuncs.com:3306)/igetcool-yuwen-online")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query(sql1)

	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// 将查询结果转为结构体数组
	var answers []Answer
	for rows.Next() {
		var answer Answer
		var instruction string
		var input string
		var output string
		var mark_num int
		err := rows.Scan(&instruction, &input, &output, &mark_num)
		if err != nil {
			panic(err.Error())
			continue
		}

		var ins, in, out []Content

		err = json.Unmarshal([]byte(instruction), &ins)
		if err != nil {
			panic(err.Error())
			continue
		}
		doc, err2 := html.Parse(strings.NewReader(ins[0].Content))
		if err2 != nil {
			panic(err2)
			continue
		}
		answer.Instruction = "例题\n" + extractText(doc)
		index := strings.Index(answer.Instruction, "备注：")
		if index != -1 {
			answer.Instruction = answer.Instruction[:index]
		}

		err = json.Unmarshal([]byte(input), &in)
		if err != nil {
			panic(err.Error())
			continue
		}
		answer.Instruction = answer.Instruction + "\n学生作答\n" + in[0].Content

		// answer.Input = in[0].Content
		answer.Input = ""

		err = json.Unmarshal([]byte(output), &out)
		if err != nil {
			panic(err.Error())
			continue
		}
		// 删除「xxx同学」中的「xxx」
		index = strings.Index(out[0].Content, "同学")
		if index != -1 {
			answer.Output = out[0].Content[index:]
		} else {
			answer.Output = out[0].Content
		}

		answer.Output = fmt.Sprintf("%d星。%s\n", mark_num, answer.Output)

		answers = append(answers, answer)

	}

	len_answers := len(answers)
	len_return := len_answers
	// if len_answers < 3000 {
	// 	fmt.Printf("len: %d => ", len_answers)
	// 	for len_answers < 3000 {
	// 		answers = append(answers, answers...)
	// 		len_answers = len(answers)
	// 	}
	// 	answers = answers[:3000]
	// 	fmt.Printf("%d => %d\n", len_answers, len(answers))

	// }

	// 将结构体数组转为JSON格式并输出到控制台
	jsonBytes, err := json.Marshal(answers)
	if err != nil {
		panic(err.Error())
	}
	return string(jsonBytes), len_return

}

func extractText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}

	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		return ""
	}

	var text string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text += extractText(c) + " "
	}

	return strings.Join(strings.Fields(text), " ")
}
