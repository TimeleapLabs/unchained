-- AlterTable
ALTER TABLE "AssetPrice" RENAME COLUMN "signature"    TO "oldSignature";
ALTER TABLE "AssetPrice"    ADD COLUMN "signature"    BYTEA;
ALTER TABLE "AssetPrice"  ALTER COLUMN "oldSignature" DROP NOT NULL;

-- AlterTable
ALTER TABLE "Signer" RENAME COLUMN "key"    TO "oldKey";
ALTER TABLE "Signer"    ADD COLUMN "key"    BYTEA;
ALTER TABLE "Signer"  ALTER COLUMN "oldKey" DROP NOT NULL;

-- CreateIndex
DROP INDEX IF EXISTS "Signer_key_key";
CREATE UNIQUE INDEX "Signer_key_key" ON "Signer"("key");
CREATE UNIQUE INDEX "Signer_oldkey_key" ON "Signer"("oldKey");
