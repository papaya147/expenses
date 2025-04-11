function openCategoryModal() {
  document.getElementById("category-modal").classList.remove("hidden");
  document.getElementById("new-category-name").focus();
}

function closeCategoryModal() {
  document.getElementById("category-modal").classList.add("hidden");
}

async function addCategory() {
  const name = document.getElementById("new-category-name").value.trim();
  if (!name) return;

  // Send POST request to create new category
  const res = await fetch("/category/create", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ name }),
  });

  if (res.ok) {
    // Optionally add it to the datalist right away
    const option = document.createElement("option");
    option.value = name;
    document.getElementById("category-options").appendChild(option);

    // Fill the input with the new category
    document.getElementById("category").value = name;

    closeCategoryModal();
  } else {
    alert("Failed to create category");
  }
}
