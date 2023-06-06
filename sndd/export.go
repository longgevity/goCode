package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	// "regexp"
)

// type User struct {
// 	Id   int
// 	Name string
// }

type Answer struct {
	Content string
	Comment string
}

type Content struct {
	Content string `json:"content"`
}

type Comment struct {
	Comment string `json:"content"`
}

func main() {
	// 连接数据库
	db, err := sql.Open("mysql", "yuwen_read:3t*hnCqIqTaE3woP@tcp(rr-2ze5mv4l0fp4c0dn3.mysql.rds.aliyuncs.com:3306)/igetcool-yuwen-online")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// 查询数据
	rows, err := db.Query(`
        SELECT ycs.comment_content content,yc.comment_content comment FROM yuwen_comment_relation ycrs 
         left join yuwen_comment ycs on ycs.comment_relation_pid = ycrs.rela_pid
         left join yuwen_comment_relation ycr on ycr.parent_id = ycrs.rela_pid
         left join yuwen_comment yc on yc.comment_relation_pid = ycr.rela_pid
         left join yuwen_schedule_work ysw on ysw.id = ycr.schedule_work_id
         left join yuwen_work yw on ysw.work_id = yw.id
         where  yw.content like "%请用自然观察法和生活观察法观察你喜欢的小动物%" and ycs.comment_content not like "%\"content\":\"https:%"
         and ycrs.identity =1 and ycrs.comment_type = 1 and yw.question_type =0
        and ycrs.mark_num = 5  limit 300
        `)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// 将查询结果转为结构体数组
	var answers []Answer
	for rows.Next() {
		var answer Answer
		// err := rows.Scan(&answer.Content, &answer.Comment)

		// if err != nil {
		// 	panic(err.Error())
		// 	// continue
		// }
		var content string
		var comment string
		err := rows.Scan(&content, &comment)
		// if err != nil {
		//  panic(err.Error())
		//  // continue
		// }
		var c []Content
		err = json.Unmarshal([]byte(content), &c)
		if err != nil {
			fmt.Println(err)
			return
		}
		answer.Content = c[0].Content

		var c2 []Comment
		err = json.Unmarshal([]byte(comment), &c2)
		if err != nil {
			fmt.Println(err)
			return
		}
		answer.Comment = c2[0].Comment

		// re := regexp.MustCompile(`\b\w+同学，\b`)
		// // str = re.ReplaceAllString(str, "同学，")

		// // 使用strings.Replace函数替换"xxx同学"
		// comment = re.ReplaceAllString(c2[0].Comment, "同学，")
		// re = regexp.MustCompile(`\b，\w+老师，\b`)
		// // str = re.ReplaceAllString(str, "同学，")

		// // 使用strings.Replace函数替换"xxx同学"
		// comment = re.ReplaceAllString(comment, "，老师")
		// answer.Comment = comment

		answers = append(answers, answer)
	}

	// 将结构体数组转为JSON格式并输出到控制台
	jsonBytes, err := json.Marshal(answers)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(string(jsonBytes))
}
