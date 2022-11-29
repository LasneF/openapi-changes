// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package tui

import (
    "fmt"
    "github.com/gdamore/tcell/v2"
    "github.com/rivo/tview"
    "what-changed/model"
)

func BuildCommitTable(commitHistory []*model.Commit) *tview.Table {
    table := tview.NewTable().SetBorders(false)
    table.SetCell(0, 0, tview.NewTableCell("Date").SetTextColor(CYAN_CELL_COLOR))
    table.SetCell(0, 1, tview.NewTableCell("Message").SetTextColor(CYAN_CELL_COLOR))
    table.SetCell(0, 2, tview.NewTableCell("Changes").SetTextColor(CYAN_CELL_COLOR))
    table.SetCell(0, 3, tview.NewTableCell("Breaking").SetTextColor(CYAN_CELL_COLOR))
    n := 1
    for k := range commitHistory {
        if commitHistory[k].Changes != nil {
            table.SetCell(n, 0, tview.NewTableCell(commitHistory[k].CommitDate.Format("01/02/06")))
            table.SetCell(n, 1, tview.NewTableCell(commitHistory[k].Message))
            if commitHistory[k].Changes != nil {
                table.SetCell(n, 2, tview.NewTableCell(fmt.Sprint(commitHistory[k].Changes.TotalChanges())).SetAlign(tview.AlignCenter))
                table.SetCell(n, 3, tview.NewTableCell(fmt.Sprint(commitHistory[k].Changes.TotalBreakingChanges())).SetAlign(tview.AlignCenter))
            } else {
                table.SetCell(n, 2, tview.NewTableCell("X").SetAlign(tview.AlignCenter))
                table.SetCell(n, 3, tview.NewTableCell("X").SetAlign(tview.AlignCenter))
            }
            n++
        }
    }

    return table
}

func RegisterModelsWithCommitTable(table *tview.Table,
    commitHistory []*model.Commit, treeView *tview.TreeView, app *tview.Application) {

    table.Select(0, 0).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
        if key == tcell.KeyEscape {
            //app.Stop()
        }
        if key == tcell.KeyEnter {
            //     table.SetSelectable(true, false)
        }
    }).SetSelectedFunc(func(row int, column int) {
        for k := 1; k < len(commitHistory); k++ {
            table.GetCell(k, 0).SetTextColor(tcell.ColorWhite)
            table.GetCell(k, 1).SetTextColor(tcell.ColorWhite)
            table.GetCell(k, 2).SetTextColor(tcell.ColorWhite)
            table.GetCell(k, 3).SetTextColor(tcell.ColorWhite)
        }
        table.GetCell(row, 0).SetTextColor(MAGENTA_CELL_COLOR)
        table.GetCell(row, 1).SetTextColor(MAGENTA_CELL_COLOR)
        table.GetCell(row, 2).SetTextColor(MAGENTA_CELL_COLOR)
        table.GetCell(row, 3).SetTextColor(MAGENTA_CELL_COLOR)

        c := commitHistory[row-1]
        activeCommit = c

        r := BuildTreeModel(c.Document, c.Changes)
        treeView.SetRoot(r)
        treeView.SetCurrentNode(r)
        app.SetFocus(treeView)
        treeView.SetTopLevel(0)
    })

}
