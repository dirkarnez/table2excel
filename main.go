// package main

// import (
// 	"fmt"
// 	"strings"

// 	"golang.org/x/net/html"
// )

// func parse(text string) (data [][]string) {

// 	tkn := html.NewTokenizer(strings.NewReader(text))

// 	var vals [][]string

// 	var isLi bool

// 	for {

// 		tt := tkn.Next()

// 		switch {

// 		case tt == html.ErrorToken:
// 			return vals

// 		case tt == html.StartTagToken:

// 			t := tkn.Token()
// 			isLi = t.Data == "li"

// 		case tt == html.TextToken:

// 			t := tkn.Token()

// 			if isLi {
// 				vals = append(vals, t.Data)
// 			}

// 			isLi = false
// 		}
// 	}
// }

// func main() {

// 	data := parse(`
// 	<table>
//   <tr>
//     <th>Company</th>
//     <th>Contact</th>
//     <th>Country</th>
//   </tr>
//   <tr>
//     <td>Alfreds Futterkiste</td>
//     <td>Maria Anders</td>
//     <td>Germany</td>
//   </tr>
//   <tr>
//     <td>Centro comercial Moctezuma</td>
//     <td>Francisco Chang</td>
//     <td>Mexico</td>
//   </tr>
//   <tr>
//     <td>Ernst Handel</td>
//     <td>Roland Mendel</td>
//     <td>Austria</td>
//   </tr>
//   <tr>
//     <td>Island Trading</td>
//     <td>Helen Bennett</td>
//     <td>UK</td>
//   </tr>
//   <tr>
//     <td>Laughing Bacchus Winecellars</td>
//     <td>Yoshi Tannamuri</td>
//     <td>Canada</td>
//   </tr>
//   <tr>
//     <td>Magazzini Alimentari Riuniti</td>
//     <td>Giovanni Rovelli</td>
//     <td>Italy</td>
//   </tr>
// </table>
// `)
// 	fmt.Println(data)
// }

package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
	"golang.org/x/net/html"
)

func main() {
	var body = strings.NewReader(`     
<html> 
  <body>
    <table>
    <tr data-height="40px">
    <td data-width="668.984px" data-has-top-border="true" data-has-left-border="true" data-has-right-border="true" data-has-bottom-border="true" data-colspan="3" data-is-bold="false">Client: Enterprise Development Holdings Limited</td>
    <td data-width="59.6562px" data-has-top-border="true" data-has-left-border="false" data-has-right-border="true" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">Ref:</td>
  </tr>
  <tr data-height="26px">
    <td data-width="668.984px" data-has-top-border="true" data-has-left-border="true" data-has-right-border="true" data-has-bottom-border="true" data-colspan="3" data-is-bold="false">Year end: 31 Dec 2021</td>
    <td data-width="60.1562px" data-has-top-border="false" data-has-left-border="true" data-has-right-border="true" data-has-bottom-border="false" data-colspan="1" data-is-bold="true">&nbsp;</td>
  </tr>
  <tr data-height="20px">
    <td data-width="668.984px" data-has-top-border="true" data-has-left-border="true" data-has-right-border="true" data-has-bottom-border="true" data-colspan="3" data-is-bold="false">File no: F912901</td>
    <td data-width="59.6562px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="true" data-has-bottom-border="true" data-colspan="1" data-is-bold="true">&nbsp;</td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="491.328px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="60.6562px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="true">&nbsp;</td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="true">AUDIT FILE INDEX</td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="491.328px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="true">&nbsp;</td>
    <td data-width="60.6562px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="491.328px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="60.6562px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">AP-A</td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="491.328px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">FINAL COMPLETION</td>
    <td data-width="60.6562px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">AP-B</td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="491.328px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">AUDIT COMPLETION</td>
    <td data-width="60.6562px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">AP-C</td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="491.328px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">AUDIT PLANNING</td>
    <td data-width="60.6562px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">AP-D</td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="491.328px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">OPTIONAL PROGRAMMES</td>
    <td data-width="60.6562px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">AP-E</td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="491.328px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">INTANGIBLE ASSETS</td>
    <td data-width="60.6562px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">AP-F</td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="491.328px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">PROPERTY, PLANT AND EQUIPMENT</td>
    <td data-width="60.6562px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">AP-G</td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="491.328px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">INVESTMENTS IN GROUP AND ASSOCIATED UNDERTAKINGS</td>
    <td data-width="60.6562px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">AP-H</td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="491.328px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">OTHER INVESTMENTS (EQUITY AND DEBT SECURITIES)</td>
    <td data-width="60.6562px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">AP-I</td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="491.328px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">INVENTORIES</td>
    <td data-width="60.6562px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">AP-J</td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="491.328px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">TRADE & OTHER RECEIVABLES & AMOUNT DUE FROM RELATED PARTIES</td>
    <td data-width="60.6562px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">AP-K</td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="491.328px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">BANK BALANCES AND CASH</td>
    <td data-width="60.6562px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">AP-L</td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="491.328px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">TRADE & OTHER PAYABLES, ACCRUALS AND AMOUNTS DUE TO RELATED PARTIES</td>
    <td data-width="60.6562px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">AP-M</td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="491.328px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">LONG-TERM LOANS AND DEFERRED INCOME</td>
    <td data-width="60.6562px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">AP-N</td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="491.328px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">PROVISIONS, CONTINGENT LIABILITIES AND FINANCIAL COMMITMENTS</td>
    <td data-width="60.6562px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">AP-O</td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="491.328px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">SHARE CAPITAL, RESERVES AND STATUTORY RECORDS</td>
    <td data-width="60.6562px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">AP-P</td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="491.328px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">INCOME TAXES</td>
    <td data-width="60.6562px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">AP-Q</td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="491.328px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">EXPENDITURES</td>
    <td data-width="60.6562px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">AP-R</td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="491.328px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">REVENUES</td>
    <td data-width="60.6562px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">AP-S</td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="491.328px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">OPERATING EFFECTIVENESS OF CONTROLS</td>
    <td data-width="60.6562px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">AP-T</td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="491.328px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">SUBSEQUENT EVENTS</td>
    <td data-width="60.6562px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">AP-V</td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="491.328px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">GENERAL LEDGER</td>
    <td data-width="60.6562px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="0px">
    <td data-width="168px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">W</td>
    <td data-width="auto" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="493.333px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">CONSOLIDATION</td>
    <td data-width="auto" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="0px">
    <td data-width="168px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">X</td>
    <td data-width="auto" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="493.333px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">ACCOUNTS WORKING PAPERS</td>
    <td data-width="auto" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="0px">
    <td data-width="168px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">Y</td>
    <td data-width="auto" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="493.333px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">OTHER PRIMARY FINANCIAL\n STATEMENTS</td>
    <td data-width="auto" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="0px">
    <td data-width="168px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">Z</td>
    <td data-width="auto" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="493.333px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">COMPUTER REPORTS AND\n RECORDS RECEIVED</td>
    <td data-width="auto" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="491.328px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="60.6562px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="0px">
    <td data-width="auto" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="true">APPENDICES</td>
    <td data-width="auto" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="493.333px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="auto" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="0px">
    <td data-width="168px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">1</td>
    <td data-width="auto" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="493.333px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">BUSINESS COMBINATIONS</td>
    <td data-width="auto" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="0px">
    <td data-width="168px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">2</td>
    <td data-width="auto" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="493.333px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">INVESTMENT PROPERTIES</td>
    <td data-width="auto" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="0px">
    <td data-width="168px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">3</td>
    <td data-width="auto" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="493.333px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">OTHER FINANCIAL\n INSTRUMENTS AND DERIVATIVES</td>
    <td data-width="auto" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="0px">
    <td data-width="168px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">4</td>
    <td data-width="auto" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="493.333px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">CONSTRUCTION CONTRACTS</td>
    <td data-width="auto" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="0px">
    <td data-width="168px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">5</td>
    <td data-width="auto" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="493.333px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">REVENUE UNDER HKFRS 15</td>
    <td data-width="auto" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="491.328px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="60.6562px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="20px">
    <td data-width="165px" data-has-top-border="true" data-has-left-border="true" data-has-right-border="true" data-has-bottom-border="true" data-colspan="1" data-is-bold="false">Client: Enterprise Developme</td>
    <td data-width="8.15625px" data-has-top-border="true" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="true" data-colspan="1" data-is-bold="false"></td>
    <td data-width="490.828px" data-has-top-border="true" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="true" data-colspan="1" data-is-bold="false"></td>
    <td data-width="59.6562px" data-has-top-border="true" data-has-left-border="true" data-has-right-border="true" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">Ref:</td>
  </tr>
  <tr data-height="20px">
    <td data-width="165px" data-has-top-border="false" data-has-left-border="true" data-has-right-border="true" data-has-bottom-border="true" data-colspan="1" data-is-bold="false">Year end: 31 Dec 2021</td>
    <td data-width="8.15625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="true" data-colspan="1" data-is-bold="false"></td>
    <td data-width="490.828px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="true" data-colspan="1" data-is-bold="false"></td>
    <td data-width="59.6562px" data-has-top-border="false" data-has-left-border="true" data-has-right-border="true" data-has-bottom-border="true" data-colspan="1" data-is-bold="true">AP-A</td>
  </tr>
  <tr data-height="20px">
    <td data-width="165px" data-has-top-border="false" data-has-left-border="true" data-has-right-border="true" data-has-bottom-border="true" data-colspan="1" data-is-bold="false">File no: F912901</td>
    <td data-width="8.15625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="true" data-colspan="1" data-is-bold="false"></td>
    <td data-width="490.828px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="true" data-colspan="1" data-is-bold="false"></td>
    <td data-width="108.156px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="491.328px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="60.6562px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="true">&nbsp;</td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="true">A</td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="491.328px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="true">FINAL COMPLETION</td>
    <td data-width="60.6562px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="491.328px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="60.6562px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="25px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">1</td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="490.828px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">Final completion memorandum</td>
    <td data-width="59.6562px" data-has-top-border="true" data-has-left-border="true" data-has-right-border="true" data-has-bottom-border="true" data-colspan="1" data-is-bold="false">AP-A1</td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="491.328px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="60.6562px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">1.2</td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="490.828px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">Final subsequent events review</td>
    <td data-width="59.6562px" data-has-top-border="true" data-has-left-border="true" data-has-right-border="true" data-has-bottom-border="true" data-colspan="1" data-is-bold="false">AP-A1.2</td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="491.328px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="60.6562px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">1.3</td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="490.828px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">Final completion checklist</td>
    <td data-width="59.6562px" data-has-top-border="true" data-has-left-border="true" data-has-right-border="true" data-has-bottom-border="true" data-colspan="1" data-is-bold="false">AP-A1.3</td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="491.328px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="60.6562px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">2</td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="491.328px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">Accounts & tax computations</td>
    <td data-width="60.6562px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="491.328px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="60.6562px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">2.1</td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="490.828px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">Final audited financial statements</td>
    <td data-width="59.6562px" data-has-top-border="true" data-has-left-border="true" data-has-right-border="true" data-has-bottom-border="true" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="491.328px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="60.6562px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="true" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">2.2</td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="490.828px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">Profits Tax computation</td>
    <td data-width="59.6562px" data-has-top-border="false" data-has-left-border="true" data-has-right-border="true" data-has-bottom-border="true" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="491.328px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="60.6562px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="true" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">3</td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="490.828px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">Adjustment Summary</td>
    <td data-width="59.6562px" data-has-top-border="true" data-has-left-border="true" data-has-right-border="true" data-has-bottom-border="true" data-colspan="1" data-is-bold="false">N/A</td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="491.328px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="60.6562px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="true" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">4</td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="490.828px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">Draft accounts, typing instructions, calling over checklist  (Aapp1)</td>
    <td data-width="59.6562px" data-has-top-border="true" data-has-left-border="true" data-has-right-border="true" data-has-bottom-border="true" data-colspan="1" data-is-bold="false">N/A</td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="491.328px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="60.6562px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="true" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">5</td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="490.828px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">Letter of representation</td>
    <td data-width="59.6562px" data-has-top-border="false" data-has-left-border="true" data-has-right-border="true" data-has-bottom-border="true" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="491.328px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="60.6562px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="true" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">6</td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="490.828px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">Letter to management</td>
    <td data-width="59.6562px" data-has-top-border="true" data-has-left-border="true" data-has-right-border="true" data-has-bottom-border="true" data-colspan="1" data-is-bold="false">N/A</td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="491.328px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="60.6562px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="true" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">7</td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="490.828px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">Company accounts disclosure checklists</td>
    <td data-width="59.6562px" data-has-top-border="true" data-has-left-border="true" data-has-right-border="true" data-has-bottom-border="true" data-colspan="1" data-is-bold="false">N/A</td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="491.328px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="60.6562px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="true" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">8.1</td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="490.828px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">Trial Balance / Management Accounts</td>
    <td data-width="59.6562px" data-has-top-border="false" data-has-left-border="true" data-has-right-border="true" data-has-bottom-border="true" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="491.328px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="60.6562px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">8.2</td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="490.828px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">Prior year's Audited Financial Statements</td>
    <td data-width="59.6562px" data-has-top-border="true" data-has-left-border="true" data-has-right-border="true" data-has-bottom-border="true" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="491.328px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="60.6562px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
  </tr>
  <tr data-height="20px">
    <td data-width="166px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">11</td>
    <td data-width="8.65625px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false"></td>
    <td data-width="490.828px" data-has-top-border="false" data-has-left-border="false" data-has-right-border="false" data-has-bottom-border="false" data-colspan="1" data-is-bold="false">Archive Form</td>
    <td data-width="59.6562px" data-has-top-border="true" data-has-left-border="true" data-has-right-border="true" data-has-bottom-border="true" data-colspan="1" data-is-bold="false">AP-A11</td>
  </tr>
    </table>
  </body>
</html>
`)
	// f := excelize.NewFile()
	// defer func() {
	// 	if err := f.Close(); err != nil {
	// 		fmt.Println(err)
	// 	}
	// }()

	// // for _, sheetName := range f.GetSheetList() {
	// // 	fmt.Println(sheetName)
	// // 	err := f.DeleteSheet(sheetName)
	// //   if err != nil {
	// //     fmt.Println(err)
	// //   }
	// // }

	// // Create a new sheet.
	// _, err := f.NewSheet("Ai")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// f.DeleteSheet("Sheet1")

	// template, err := excelize.OpenFile("AP2.2 (2023.04.01).xlsx")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer func() {
	// 	// Close the spreadsheet.
	// 	if err := template.Close(); err != nil {
	// 		fmt.Println(err)
	// 	}
	// }()

	// rows, err := template.GetRows("Ai")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("len(rows)", len(rows))
	// for i, row := range rows {
	// 	cell, _ := excelize.CoordinatesToCellName(1, i+1) //A1A2A3A4...

	// 	_, rowNumeric, _ := excelize.SplitCellName(cell)

	// 	rowHeight, _ := template.GetRowHeight("Ai", rowNumeric)
	// 	fmt.Println("len(row)", len(row), "!!!!", cell, "***", rowHeight)
	// 	f.SetRowHeight("Ai", rowNumeric, rowHeight)

	// 	f.SetSheetRow("Ai", cell, &row)

	// 	for j, colCell := range row {
	// 		sss, _ := excelize.CoordinatesToCellName(j+1, i+1)
	// 		colAlpha, _, _ := excelize.SplitCellName(sss)
	// 		colWidth, _ := template.GetColWidth("Ai", colAlpha)
	// 		f.SetColWidth("Ai", colAlpha, colAlpha, colWidth)
	// 		fmt.Println("????", sss)
	// 		fmt.Print(colCell, "\t")
	// 	}

	// 	fmt.Println()
	// }

	// mergeCells, _ := template.GetMergeCells("Ai")
	// for _, mergeCell := range mergeCells {
	// 	mergeCellCoord := strings.Split(mergeCell[0], ":")
	// 	fmt.Println("Ai", mergeCellCoord[0], mergeCellCoord[1])
	// 	f.MergeCell("Ai", mergeCellCoord[0], mergeCellCoord[1])
	// }

	// // aw, _ := template.GetColWidth("Ai", "A")
	// // f.SetColWidth("Ai", "A", "A", aw)

	// // bw, _ := template.GetColWidth("Ai", "B")
	// // f.SetColWidth("Ai", "B", "B", bw)

	// // cw, _ := template.GetColWidth("Ai", "C")
	// // f.SetColWidth("Ai", "C", "C", cw)

	// // dw, _ := template.GetColWidth("Ai", "D")
	// // f.SetColWidth("Ai", "D", "D", dw)

	// templatePageLayout, _ := template.GetPageLayout("Ai")
	// f.SetPageLayout("Ai", &templatePageLayout)

	// templatePageMargins, _ := template.GetPageMargins("Ai")
	// f.SetPageMargins("Ai", &templatePageMargins)

	// templateSheetProps, _ := template.GetSheetProps("Ai")
	// f.SetSheetProps("Ai", &templateSheetProps)

	// templateSheetVisible, _ := template.GetSheetVisible("Ai")
	// f.SetSheetVisible("Ai", templateSheetVisible)

	// templateDefaultFont, _ := template.GetDefaultFont()
	// f.SetDefaultFont(templateDefaultFont)

	// templateSheetDimension, _ := template.GetSheetDimension("Ai")
	// f.SetSheetDimension("Ai", templateSheetDimension)

	// // Save spreadsheet by the given path.
	// if err := f.SaveAs("Book1.xlsx"); err != nil {
	// 	fmt.Println(err)
	// }

	// for _, col := range cols {
	// 	for _, rowCell := range col {
	// 		fmt.Print(rowCell, "\t")
	// 	}
	// 	fmt.Println()
	// }

	// f.GetColWidth("Ai", "B2")
	// // Get value from cell by given worksheet name and cell reference.
	// cell, err := f.GetCellValue
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(cell)
	// // Get all the rows in the Sheet1.
	// rows, err := f.GetRows("Sheet1")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// for _, row := range rows {
	// 	for _, colCell := range row {
	// 		fmt.Print(colCell, "\t")
	// 	}
	// 	fmt.Println()
	// }

	//temp, _ := excelize.OpenFile("Test.xlsx")

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	f.SetColWidth("Sheet1", "A", "A", 23.29)
	f.SetColWidth("Sheet1", "B", "B", 0.83)
	f.SetColWidth("Sheet1", "C", "C", 69.71)
	f.SetColWidth("Sheet1", "D", "D", 8.29)

	// for _, sheetName := range f.GetSheetList() {
	// 	fmt.Println(sheetName)
	// 	err := f.DeleteSheet(sheetName)
	//   if err != nil {
	//     fmt.Println(err)
	//   }
	// }

	// indexOf := func(element string) int {
	// 	for k, v := range []string{
	// 		"none",
	// 		"thin",
	// 		"medium",
	// 		"dashed",
	// 		"dotted",
	// 		"thick",
	// 		"double",
	// 		"hair",
	// 		"mediumDashed",
	// 		"dashDot",
	// 		"mediumDashDot",
	// 		"dashDotDot",
	// 		"mediumDashDotDot",
	// 		"slantDashDot",
	// 	} {
	// 		if element == v {
	// 			return k
	// 		}
	// 	}
	// 	return -1 //not found.
	// }

	// for _, border := range temp.Styles.Borders.Border[1:] {
	// 	//fmt.Println(border.Top.Color.Auto, border.Left.Color, border.Right.Color, border.Bottom.Color, border.Diagonal.Color)

	// 	left := excelize.Border{Type: "left", Style: indexOf(border.Left.Style)}
	// 	if border.Left.Color == nil || border.Left.Color.Auto {
	// 		left.Color = "#000000"
	// 	}

	// 	top := excelize.Border{Type: "top", Style: indexOf(border.Top.Style)}
	// 	if border.Top.Color == nil || border.Top.Color.Auto {
	// 		top.Color = "#000000"
	// 	}

	// 	bottom := excelize.Border{Type: "bottom", Style: indexOf(border.Bottom.Style)}
	// 	if border.Bottom.Color == nil || border.Bottom.Color.Auto {
	// 		bottom.Color = "#000000"
	// 	}

	// 	right := excelize.Border{Type: "right", Style: indexOf(border.Right.Style)}
	// 	if border.Right.Color == nil || border.Right.Color.Auto {
	// 		right.Color = "#000000"
	// 	}

	// 	f.NewStyle(&excelize.Style{
	// 		Border: []excelize.Border{left, top, bottom, right},
	// 		//Fill: excelize.Fill{}
	// 	})

	// 	fmt.Println("====================")
	// }

	// // for _, fill := range temp.Styles.Fills.Fill[1:] {
	// // 	fmt.Println(fill

	// // }
	// fmt.Println("!!!!!!!!!!!!!!!!!")

	// ReadArea(temp, "Sheet1", 9, 6, 20, 20, func(row, col int, cellName, value string) {
	// 	cellStyles, _ := temp.GetCellStyle("Sheet1", cellName)
	// 	f.SetCellStyle("Sheet1", cellName, cellName, cellStyles)
	// 	fmt.Println(cellName, cellStyles)
	// })

	// styleBold, _ := f.NewStyle(&excelize.Style{
	// 	Font: &excelize.Font{
	// 		Bold:   true,
	// 		Family: "Arial",
	// 		Size:   9,
	// 		Color:  "000000",
	// 	},
	// 	Alignment: &excelize.Alignment{
	// 		Horizontal: "left",
	// 		Vertical:   "top",
	// 	},
	// })

	// styleC, _ := f.NewStyle(&excelize.Style{
	// 	Border: []excelize.Border{
	// 		{Type: "left", Color: "000000", Style: 1},
	// 		{Type: "top", Color: "000000", Style: 1},
	// 		{Type: "bottom", Color: "000000", Style: 1},
	// 		{Type: "right", Color: "000000", Style: 1},
	// 	},
	// 	Font: &excelize.Font{
	// 		Family: "Arial",
	// 		Size:   9,
	// 		Color:  "000000",
	// 	},
	// 	Alignment: &excelize.Alignment{
	// 		Horizontal: "left",
	// 	},
	// })
	HTMLTableToExcel(body, func(rowIndex int, height float64) {
		fmt.Println("onRow", rowIndex, height)
		f.SetRowHeight("Sheet1", rowIndex+1, height)
	}, func(rowIndex, cellIndex, colSpan int, width float64, hasTopBorder, hasLeftBorder, hasRightBorder, hasBottomBorder, isBold bool, value string) {
		cellName, _ := excelize.CoordinatesToCellName(cellIndex+1, rowIndex+1)

		//ascii := string(rune(+65))
		//fmt.Println("!!!!", rowIndex, cellIndex, width)
		// if rowIndex == 5 {
		// 	ascii := string(rune(cellIndex + 65))
		// 	fmt.Println("!", ascii, width)
		// 	f.SetColWidth("Sheet1", ascii, ascii, width)
		// }

		// if hasBorder {
		// 	f.SetCellStyle("Sheet1", cellName, cellName, styleC)
		// } else if isBold {
		// 	f.SetCellStyle("Sheet1", cellName, cellName, styleBold)
		// } else {
		// 	f.SetCellStyle("Sheet1", cellName, cellName, styleGlobal)
		// }
		// {Type: "left", Color: "000000", Style: 1},
		// {Type: "top", Color: "000000", Style: 1},
		// {Type: "bottom", Color: "000000", Style: 1},
		// {Type: "right", Color: "000000", Style: 1},

		style := excelize.Style{
			Font: &excelize.Font{
				Family: "Arial",
				Size:   9,
				Color:  "000000",
			},
			Alignment: &excelize.Alignment{
				Horizontal: "left",
				Vertical:   "top",
			},
		}

		if hasTopBorder {
			style.Border = append(style.Border, excelize.Border{Type: "top", Color: "000000", Style: 1})
		}

		if hasLeftBorder {
			style.Border = append(style.Border, excelize.Border{Type: "left", Color: "000000", Style: 1})
		}

		if hasBottomBorder {
			style.Border = append(style.Border, excelize.Border{Type: "bottom", Color: "000000", Style: 1})
		}

		if hasRightBorder {
			style.Border = append(style.Border, excelize.Border{Type: "right", Color: "000000", Style: 1})
		}

		style.Font.Bold = isBold

		styleID, err := f.NewStyle(&style)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("styleID", styleID, cellName)

		if colSpan > 1 {
			endCellName, _ := excelize.CoordinatesToCellName(cellIndex+colSpan, rowIndex+1)
			f.MergeCell("Sheet1", cellName, endCellName)
			f.SetCellStyle("Sheet1", cellName, endCellName, styleID)
		} else {
			f.SetCellStyle("Sheet1", cellName, cellName, styleID)
		}

		f.SetCellValue("Sheet1", cellName, value)
	})

	//f.MergeCell("Sheet1", "A1", "B1")

	// styleA, _ := f.NewStyle(&excelize.Style{
	// 	Border: []excelize.Border{
	// 		{Type: "left", Color: "000000", Style: 1},
	// 		{Type: "top", Color: "000000", Style: 1},
	// 		{Type: "bottom", Color: "000000", Style: 1},
	// 		{Type: "right", Color: "000000", Style: 1},
	// 	},
	// })

	// styleB, _ := f.NewStyle(&excelize.Style{
	// 	Font: &excelize.Font{
	// 		Bold:   false,
	// 		Family: "Arial",
	// 		Size:   10,
	// 		Color:  "000000",
	// 	},
	// })

	//f.SetCellValue("Sheet1", "A1", 1234)

	if err := f.SaveAs("OUTPUT.xlsx"); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Done")
}

func ReadArea(target *excelize.File, sheet string, startingRow, startingCol, endingRowExclusive, endingColExclusive int, cellCallback func(row, col int, cellName, value string)) {
	for row := startingRow; row < endingRowExclusive; row++ {
		for col := startingCol; col < endingColExclusive; col++ {
			cellName, _ := excelize.CoordinatesToCellName(col, row)
			val, _ := target.GetCellValue(sheet, cellName)
			cellCallback(row, col, cellName, val)
		}
	}
}

func HTMLTableToExcel(body *strings.Reader, onRow func(rowIndex int, height float64), onCell func(rowIndex, cellIndex, colSpan int, width float64, hasTopBorder, hasLeftBorder, hasRightBorder, hasBottomBorder, isBold bool, value string)) {
	z := html.NewTokenizer(body)
	trIndex := -1
	tdIndex := -1

	getTrueFalse := func(attrMap map[string]string, key string) bool {
		v, ok := attrMap[key]
		if !ok {
			return false
		}

		if v == "true" {
			return true
		}
		return false
	}

	for z.Token().Data != "html" {
		tt := z.Next()
		if tt == html.StartTagToken {
			t := z.Token()

			if t.Data == "tr" {
				tdIndex = -1
				trIndex = trIndex + 1
				// if len(row) > 0 {
				// 	table = append(table, row)
				// 	row = []string{}
				// }
				trAttriMap := map[string]string{}
				for _, attr := range t.Attr {
					trAttriMap[attr.Key] = attr.Val
				}

				rowHeight, ok := trAttriMap["data-height"]
				if !ok {
					rowHeight = "0px"
				}

				rowHeightPt, _ := strconv.ParseFloat(strings.TrimRight(rowHeight, "px"), 64)
				onRow(trIndex, rowHeightPt/1.3333)
			}

			if t.Data == "td" {
				tdIndex = tdIndex + 1
				inner := z.Next()

				if inner == html.TextToken {
					text := (string)(z.Text())
					//fmt.Print(t.Data, "-")

					tdAttriMap := map[string]string{}
					for _, attr := range t.Attr {
						tdAttriMap[attr.Key] = attr.Val
					}

					colWidth, ok := tdAttriMap["data-width"]
					if !ok {
						colWidth = "0px"
					}

					colWidthPt, _ := strconv.ParseFloat(strings.TrimRight(colWidth, "px"), 64)

					colSpan, ok := tdAttriMap["data-colspan"]
					if !ok {
						colSpan = "1"
					}

					colSpanInt, _ := strconv.Atoi(colSpan)

					onCell(trIndex, tdIndex,
						colSpanInt,
						colWidthPt/1.3333,
						getTrueFalse(tdAttriMap, "data-has-top-border"),
						getTrueFalse(tdAttriMap, "data-has-left-border"),
						getTrueFalse(tdAttriMap, "data-has-right-border"),
						getTrueFalse(tdAttriMap, "data-has-bottom-border"),
						getTrueFalse(tdAttriMap, "data-is-bold"),
						text,
					)

					tdIndex = tdIndex + colSpanInt - 1
					fmt.Println(tdIndex)
				}
			}
		}
	}
}
