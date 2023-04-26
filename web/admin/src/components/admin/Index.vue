<template>
  <div style="width: 300px; border: 1px solid #d9d9d9; border-radius: 4px">

    <a-calendar :fullscreen="false" :header-render="headerRender" @panelChange="onPanelChange" />

    <div style="margin-top: 20px; font-size: 15px;">
      系统成功运行：
      <a-progress type="circle" :percent="1" :format="percent => `${percent} Days`" />
    </div>

  </div>

</template>

<style scoped>
div.ant-progress-circle,
div.ant-progress-line {
  margin-right: 8px;
  margin-bottom: 8px;
}
</style>

<script>
export default {
  methods: {
    onPanelChange(value, mode) {
      // eslint-disable-next-line no-console
      console.log(value, mode)
    },
    headerRender({ value, type, onChange, onTypeChange }) {
      const start = 0
      const end = 12
      const monthOptions = []

      const current = value.clone()
      const localeData = value.localeData()
      const months = []
      for (let i = 0; i < 12; i++) {
        current.month(i)
        months.push(localeData.monthsShort(current))
      }

      for (let index = start; index < end; index++) {
        monthOptions.push(
          <a-select-option class="month-item" key={`${index}`}>
            {months[index]}
          </a-select-option>
        )
      }
      const month = value.month()

      const year = value.year()
      const options = []
      for (let i = year - 10; i < year + 10; i += 1) {
        options.push(
          <a-select-option key={i} value={i} class="year-item">
            {i}
          </a-select-option>
        )
      }
      return (
        <div>
          <div>Simplicity系统日历</div>
          <a-row type="flex" justify="space-between">
            <a-col>
              <a-radio-group size="small" onChange={(e) => onTypeChange(e.target.value)} value={type}>
                <a-radio-button value="month">Month</a-radio-button>
                <a-radio-button value="year">Year</a-radio-button>
              </a-radio-group>
            </a-col>
            <a-col>
              <a-select
                size="small"
                dropdownMatchSelectWidth={false}
                class="my-year-select"
                onChange={(newYear) => {
                  const now = value.clone().year(newYear)
                  onChange(now)
                }}
                value={String(year)}
              >
                {options}
              </a-select>
            </a-col>
            <a-col>
              <a-select
                size="small"
                dropdownMatchSelectWidth={false}
                value={String(month)}
                onChange={(selectedMonth) => {
                  const newValue = value.clone()
                  newValue.month(parseInt(selectedMonth, 10))
                  onChange(newValue)
                }}
              >
                {monthOptions}
              </a-select>
            </a-col>
          </a-row>
        </div>
      )
    },
  },
}
</script>
