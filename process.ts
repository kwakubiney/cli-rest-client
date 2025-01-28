import { execa } from "execa";
import { program } from "commander";

program
  .command("create")
  .description("Create a new project from template")
  .argument("<template>", "template name to use")
  .action(async (template) => {
    try {
      console.log(`Creating new project using template: ${template}`);
      
      const cmdOptions = [
        "new",
        "install",
        template,
        "--nuget-source",
        "https://pkgs.dev.azure.com/hubtel/_packaging/hubtel/nuget/v3/index.json",
        "--interactive"
      ];

      await execa("dotnet", cmdOptions, { stdio: "inherit" });
      console.log("Project creation completed successfully");
    } catch (error) {
      console.error("Error creating project:", error);
    }
  });

program.parse(process.argv);