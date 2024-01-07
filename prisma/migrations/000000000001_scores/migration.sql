-- AlterTable
ALTER TABLE "Signer" ADD COLUMN     "points" INTEGER NOT NULL DEFAULT 0;

-- CreateTable
CREATE TABLE "SprintPoint" (
    "sprint" INTEGER NOT NULL,
    "score" INTEGER NOT NULL,
    "signerId" INTEGER NOT NULL,

    CONSTRAINT "SprintPoint_pkey" PRIMARY KEY ("signerId","sprint")
);

-- AddForeignKey
ALTER TABLE "SprintPoint" ADD CONSTRAINT "SprintPoint_signerId_fkey" FOREIGN KEY ("signerId") REFERENCES "Signer"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

