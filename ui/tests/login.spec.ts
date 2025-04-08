import { test, expect } from "@playwright/test";

test("to look ok", async ({ page }) => {
  await page.goto("./login");
  await expect(page).toHaveScreenshot();
});
