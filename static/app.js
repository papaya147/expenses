document.addEventListener("DOMContentLoaded", function () {
  console.log("Expense Tracker App initialized");

  // Initialize the chart if canvas exists
  const chartCanvas = document.getElementById("expenseChart");
  if (chartCanvas) {
    initializeChart();
  }

  // Set default date and time if the datetime field exists
  const datetimeField = document.getElementById("datetime");
  if (datetimeField) {
    const now = new Date();
    const year = now.getFullYear();
    const month = String(now.getMonth() + 1).padStart(2, "0");
    const day = String(now.getDate()).padStart(2, "0");
    const hours = String(now.getHours()).padStart(2, "0");
    const minutes = String(now.getMinutes()).padStart(2, "0");

    datetimeField.value = `${year}-${month}-${day}T${hours}:${minutes}`;
  }
});

let expenseChart = null;

function initializeChart() {
  if (expenseChart) {
    expenseChart.destroy();
    expenseChart = null;
  }

  // Fetch chart data from the server
  fetch("/txn/chart")
    .then((response) => {
      if (!response.ok) {
        throw new Error("Failed to fetch chart data");
      }
      return response.json();
    })
    .then((data) => {
      renderChart(data);
    })
    .catch((error) => {
      console.error("Error fetching chart data:", error);

      // Show error message in the chart container
      const chartContainer = document.querySelector(".chart-container");
      if (chartContainer) {
        chartContainer.innerHTML =
          '<p class="code-font" style="text-align: center; padding: 40px 0; color: #d32f2f;">Failed to load chart data</p>';
      }
    });
}

function renderChart(data) {
  const ctx = document.getElementById("expenseChart").getContext("2d");

  // Create gradient for the background
  const gradient = ctx.createLinearGradient(0, 0, 0, 300);
  gradient.addColorStop(0, "rgba(75, 192, 192, 0.2)");
  gradient.addColorStop(1, "rgba(75, 192, 192, 0)");

  expenseChart = new Chart(ctx, {
    type: "line",
    data: {
      labels: data.timestamps,
      datasets: [
        {
          label: "Daily Net Flow (₹)",
          data: data.amounts,
          backgroundColor: gradient,
          borderColor: "rgba(75, 192, 192, 1)",
          borderWidth: 2,
          pointBackgroundColor: "rgba(75, 192, 192, 1)",
          pointBorderColor: "#fff",
          pointBorderWidth: 1,
          pointRadius: 4,
          tension: 0.3,
        },
      ],
    },
    options: {
      responsive: true,
      maintainAspectRatio: false,
      plugins: {
        legend: {
          display: false,
        },
        tooltip: {
          callbacks: {
            label: function (context) {
              let label = context.dataset.label || "";
              if (label) {
                label += ": ";
              }
              label += "₹" + parseFloat(context.raw).toFixed(2);
              return label;
            },
          },
        },
      },
      scales: {
        y: {
          grid: {
            color: "rgba(200, 200, 200, 0.1)",
          },
          ticks: {
            callback: function (value) {
              return "₹" + value;
            },
            font: {
              family: "JetBrains Mono, monospace",
            },
          },
        },
        x: {
          grid: {
            display: false,
          },
          ticks: {
            font: {
              family: "JetBrains Mono, monospace",
            },
          },
        },
      },
    },
  });
}
