
class chart {
  constructor(json){
    console.log(this.json)
  }
  drowChart(json){
    console.log(json[2].Value)
    var ctx = document.getElementById("myChart");
    var myChart = new Chart(ctx, {
      type: 'line',
      data: {
      labels: ['8月1日', '8月2日', '8月3日', '8月4日', '8月5日', '8月6日', '8月7日'],
      datasets: [
        {
          label: '最高気温(度）',
          data: json[2].Value,
          borderColor: "rgba(255,0,0,1)",
          backgroundColor: "rgba(0,0,0,0)"
        },
        {
          label: '最低気温(度）',
          data: [25, 27, 27, 25, 26, 27, 25, 21],
          borderColor: "rgba(0,0,255,1)",
          backgroundColor: "rgba(0,0,0,0)"
        }
      ],
    },
    options: {
      title: {
        display: true,
        text: '気温（8月1日~8月7日）'
      },
      scales: {
        yAxes: [{
          ticks: {
            suggestedMax: 40,
            suggestedMin: 0,
            stepSize: 10,
            callback: function(value, index, values){
              return  value +  '度'
            }
          }
        }]
      },
    }
    });
  }
}
const chartClass = new chart();
$.ajax({
  url: '/items',
  type: "get",
  dataType: 'json',
}).then(function (json) {
  console.log(json)
  chartClass.drowChart(json);
});