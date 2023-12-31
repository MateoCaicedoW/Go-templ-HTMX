package tasks

import (
	"templ/internal/tasks"
	"templ/render"

	tasksC "templ/components/tasks"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Get(c echo.Context) error {
	index := c.Request().URL.Query().Get("index")
	id, _ := uuid.Parse(c.Param("id"))
	tsk, err := tasks.Find(h.DB, id)
	if err != nil {
		return err
	}

	return render.Render(c, tasksC.Edit(tsk, index))
}
