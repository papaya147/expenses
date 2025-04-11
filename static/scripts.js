async function fetchCategories(name = "") {
  try {
    const res = await fetch("/category/list", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ name }),
    });

    if (!res.ok) throw new Error("Failed to load categories");

    return await res.json();
  } catch (err) {
    console.error("Error fetching categories:", err);
    return [];
  }
}

async function updateCategorySuggestions(name) {
  const categories = await fetchCategories(name);
  const datalist = document.getElementById("category-options");
  datalist.innerHTML = "";

  categories.forEach((cat) => {
    if (cat.name) {
      const option = document.createElement("option");
      option.value = cat.name;
      datalist.appendChild(option);
    }
  });
}

document.addEventListener("DOMContentLoaded", () => {
  const categoryInput = document.getElementById("category");

  // Initial load (empty search)
  updateCategorySuggestions("");

  // Fetch suggestions as user types (debounced)
  let timeoutId;
  categoryInput.addEventListener("input", (e) => {
    clearTimeout(timeoutId);
    timeoutId = setTimeout(() => {
      const query = e.target.value.trim();
      updateCategorySuggestions(query);
    }, 200); // Debounce: wait 200ms after user stops typing
  });
});
