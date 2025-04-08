import { test, expect } from "@playwright/test";

test("to look ok", async ({ page }) => {
  await page.goto("./login");
  await page.screenshot({ fullPage: true });
  await expect(page).toHaveScreenshot()
});
